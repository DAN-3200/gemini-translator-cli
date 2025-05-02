package agentAI

import (
	"app/pkg/utils"
	"context"
	"fmt"

	Gemini "github.com/google/generative-ai-go/genai"

	"google.golang.org/api/option"
)

const geminiKEY string = "AIzaSyCbtBOlwU1yI7BYxkanm2SPjYykkgh4xnQ"

var (
	ctx    = context.Background()
	client = InitGemini(geminiKEY)
	model  = client.GenerativeModel("gemini-1.5-flash")
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

func UseDictionary(word string) string {
	if len(word) <= 0 {
		return "[Error]"
	}

	command := `
		Translate the word ´%s´ exactly and directly. Do not ask for context or clarification.

		Format:
		class:
		[part of speech]

		meaning:
		[concise English definition, line breaks every ~20 characters]

		translation:
		[short Portuguese translation, line breaks every ~20 characters]
	`

	prompt := Gemini.Text(
		fmt.Sprintf(command, word),
	)

	response, err := model.GenerateContent(ctx, prompt)
	if err != nil {
		return "Erro na requisição de conteudo"

	}
	var description = fmt.Sprintf("%s", response.Candidates[0].Content.Parts[0])

	return description
}
