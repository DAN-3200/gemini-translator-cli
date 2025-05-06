package agentAI

import (
	"app/internal/models"
	"app/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	Gemini "github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
)

const geminiKEY string = "AIzaSyCbtBOlwU1yI7BYxkanm2SPjYykkgh4xnQ"

var (
	ctx    = context.Background()
	client = InitGemini(geminiKEY)
	model  = client.GenerativeModel("gemini-2.0-flash-lite")
)

func InitGemini(Key string) *Gemini.Client {
	client, err := Gemini.NewClient(ctx, option.WithAPIKey(geminiKEY))
	if err != nil {
		return nil
	}
	return client
}

func UseTranslation(text string, switchLang bool) string {
	langs := utils.Ternary(switchLang, "PT-BR to EN", "EN to PT-BR")
	command := "Translate from %s: '%s'. Return only the exact translation, no explanation."
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
		}
	}

	command := `
		Task: Generate a single JSON object for the word "%s".

		Requirements:
		- Do not include code blocks, backticks, or explanations.
		- Return only the JSON object.
		- Generate the phrase without a period at the end.

		Format:
		{
			"word": string,
			"partOfSpeech": string, // If it's a verb, specify the verb tense. modelo: verb - [verb tense]
			"definition": string,
			"example": string,
			"translation": string, // Translation in Portuguese (pt-br)
			"synonyms": string // separated by commas
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
	var cleanJSON = strings.ReplaceAll(dataJSON, "`", "")
	cleanJSON = strings.ReplaceAll(cleanJSON, "json", "")

	var data models.DictionaryEntry
	if err := json.Unmarshal([]byte(cleanJSON), &data); err != nil {
		return models.DictionaryEntry{Word: "Error Json.Unmarshal"}
	}

	return data
}
