package layout

import (
	. "app/pkg/design"
	u "app/pkg/utils"

	css "github.com/charmbracelet/lipgloss"
)

// -- Styles and Components -----------------------------------------------------
const (
	blue   = "#87CEEB" // #87CEEB
	green  = "#50C878" // #50C878
	yellow = "#FADA5E" // #FADA5E
	white  = "#FAF9F6" // #FAF9F6
)

// LoroApp.View() string: renderiza a interface a partir do dados model
func (it LoroApp) View() string {
	var titleApp = css.NewStyle().
		Bold(true).
		Padding(0, 2).
		Foreground(css.Color("#28282b")).
		Background(css.Color(green)).
		Render("Translate")

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
				u.Ternary(it.switchLang, styleFont("PT > EN"), styleFont("EN > PT")),
			),
		)
	}

	var textTranslate = css.NewStyle().
		Foreground(css.Color(blue)).
		Height(2).
		Render(it.textTranslate)

	var prefixStyle = func(text string) string {
		return css.NewStyle().
			Foreground(css.Color("#606060")).
			Render(text)
	}

	var textInput = func() string {
		it.textInput.Prompt = prefixStyle("From ")
		it.textInput.TextStyle = css.NewStyle().
			Foreground(css.Color(yellow))

		return it.textInput.View()
	}

	var infoHelp = css.NewStyle().
		Foreground(css.Color("#606060")).
		Render("[Esc] Exit [Ctrl+d] Switch language [Ctrl+x] Clear")

	var box = css.NewStyle().Padding(1)

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
