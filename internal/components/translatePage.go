package components

import (
	. "app/pkg/design"
	u "app/pkg/utils"
	"app/internal/types"

	css "github.com/charmbracelet/lipgloss"
)

func TranslatePage(it types.CtxMain) string {
	var langSwap = func() string {
		var styleFont = func(text string) string {
			return css.NewStyle().
				Bold(true).
				Render(text)
		}

		var styleBox = css.NewStyle().
			Padding(0, 2).
			Background(css.Color("#28282b"))

		return styleBox.Render(
			Div(
				u.Ternary(it.CtxTranslate.SwitchLang, styleFont("PT > EN"), styleFont("EN > PT")),
			),
		)
	}

	var textTranslate = css.NewStyle().
		Foreground(css.Color(blue)).
		Height(2).
		Render(it.CtxTranslate.Text)

	var textInput = func() string {
		it.TextInput.Prompt = prefixStyle("From ")
		it.TextInput.TextStyle = css.NewStyle().
			Foreground(css.Color(yellow))

		return it.TextInput.View()
	}

	return box.Render(
		Div(
			titleApp+langSwap(),
			"",
			textInput(),
			"",
			prefixStyle("To ")+textTranslate,
			infoHelp,
		),
	)
}
