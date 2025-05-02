package layout

import (
	c "app/ctx_share"

	"github.com/charmbracelet/bubbles/textinput"
)

// Contexto Geral da aplicação responsável variáveis persistentes durante a renderização
type App struct {
	c.CtxMain
}

func SetApp() App {
	var ti = textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	return App{
		CtxMain: c.CtxMain{
			TextInput:  ti,
			SwitchMode: false,
			CtxTranslate: c.CtxTranslate{
				Text:       "...",
				SwitchLang: false,
			},
			CtxDict: c.CtxDict{
				Response: "...",
			},
		},
	}
}
