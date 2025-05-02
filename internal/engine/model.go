package engine

import (
	t "app/internal/types"

	"github.com/charmbracelet/bubbles/textinput"
)

// Contexto Geral da aplicação responsável variáveis persistentes durante a renderização
type App struct {
	t.CtxMain
}

func SetApp() *App {
	var ti = textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	return &App{
		CtxMain: t.CtxMain{
			TextInput:  ti,
			SwitchMode: false,
			CtxTranslate: t.CtxTranslate{
				Text:       "...",
				SwitchLang: false,
			},
			CtxDict: t.CtxDict{
				Response: "...",
			},
		},
	}
}
