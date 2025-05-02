package layout

import (
	. "app/pkg/design"
	"app/pkg/utils"

	css "github.com/charmbracelet/lipgloss"
)

func (it App) DictPage() string {
	var dictNotification = css.NewStyle().
		Bold(true).
		Padding(0, 2).
		Background(css.Color("#28282b")).
		Render("Dictionary")

	var textInput = func() string {
		it.TextInput.Prompt = prefixStyle("Word: ")
		it.TextInput.TextStyle = css.NewStyle().
			Foreground(css.Color(yellow))

		return it.TextInput.View()
	}
	var _ = func(element string) string {
		ft := css.NewStyle().
			Padding(1, 2).
			AlignVertical(css.Position(css.Center)).
			Foreground(css.Color(blue)).
			Background(css.Color("#28282b"))

		return ft.Render(element)
	}

	return box.Render(
		Div(
			titleApp+dictNotification,
			"",
			textInput(),
			"",
			utils.Ternary(len(it.CtxDict.Response) > 0, fontColor(it.CtxDict.Response, blue), ""),
			"",
			infoHelp,
		),
	)

}
