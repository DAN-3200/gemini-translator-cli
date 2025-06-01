package agentAI

import (
	"app/internal/models"
	"app/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	Gemini "github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
)

var (
	ctx    = context.Background()
	model *Gemini.GenerativeModel
)

func InitGemini() {
	client, _ := Gemini.NewClient(ctx, option.WithAPIKey(os.Getenv("KEY")))
	model = client.GenerativeModel("gemini-2.0-flash-lite")
}

func UseTranslation(text string, switchLang bool) string {
	if text == "" {
		return "..."
	}
	langs := utils.Ternary(switchLang, "PT-BR to EN", "EN to PT-BR")
	command := `Translate from %s: "%s". Return only the exact translation, no explanation.`
	prompt := Gemini.Text(
		fmt.Sprintf(command, langs, text),
	)

	response, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return "Erro na requisição de conteudo"
	}

	var translate = fmt.Sprintf("%s", response.Candidates[0].Content.Parts[0])

	return translate
}

func UseDictionary(word string) models.DictionaryEntry {
	if len(word) <= 0 {
		return models.DictionaryEntry{
			Word:         "Word",
			PartOfSpeech: "Part of speech",
			Definition:   "Definition",
			Example:      "Example",
			Synonyms:     "Synonyms",
			Collocations: "Collocations",
		}
	}

	command := `
		Task: Return a strict JSON object for the word "%s".

		Rules:
		- Output must be only the JSON object. No text, no comments, no formatting.
		- The JSON must be valid: use double quotes and no trailing commas.
		- Use simple, learner-friendly English in all fields.
		- The example sentence must be natural, common, and end without a period.
	 	- If the word is a verb, format partOfSpeech as: "verb - [tense]".

		Schema:
		{
			"word": string,
			"partOfSpeech": string,
			"definition": string, 
			"example": string,
			"synonyms": string // comma-separated
			"collocations": string, // comma-separated; frequent word combinations 
		}
	`
	prompt := Gemini.Text(
		fmt.Sprintf(command, word),
	)

	response, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return models.DictionaryEntry{}
	}

	var dataJSON = fmt.Sprintf("%s", response.Candidates[0].Content.Parts[0])
	// [ ] Otimizar
	var cleanJSON = strings.ReplaceAll(dataJSON, "`", "")
	cleanJSON = strings.ReplaceAll(cleanJSON, "json", "")

	var data models.DictionaryEntry
	if err := json.Unmarshal([]byte(cleanJSON), &data); err != nil {
		return models.DictionaryEntry{Word: "Error Json.Unmarshal"}
	}

	return data
}
