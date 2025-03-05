package translate

import (
	"app/pkg"
	"context"
	"fmt"

	Gemini "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func UseGemini(text string, switchLang bool) string {

	// Chame Disponibilizada no IA Google Studio
	var GeminiKEY = "AIzaSyCbtBOlwU1yI7BYxkanm2SPjYykkgh4xnQ"

	// Acesso a API Gemini
	var ctx = context.Background()
	var client, err = Gemini.NewClient(ctx, option.WithAPIKey(GeminiKEY))
	if err != nil {
		fmt.Printf("Erro: %v", err)
	}
	defer client.Close()

	// Modelo generativo da IA
	var model = client.GenerativeModel("gemini-1.5-flash")

	// Geração de prompt
	var langs = pkg.Ternary(switchLang, "PT-BR to EN", "EN to PT-BR")
	var command = "Translate %s precisely: '%s'. Answer only the translation. No explanations."
	var prompt = Gemini.Text(
		fmt.Sprintf(command, langs, text),
	)

	// Requisição
	var response, err0 = model.GenerateContent(ctx, prompt)
	if err0 != nil {
		fmt.Printf("Erro na Geração de Conteudo: %v", err)
	}

	var translate = fmt.Sprintf("%s", response.Candidates[0].Content.Parts[0])

	return translate
}
