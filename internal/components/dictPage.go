package components

import (
	"app/internal/models"
	. "app/pkg/design"

	css "github.com/charmbracelet/lipgloss"
)

func DictPage(it models.CtxMain) string {
	var dictFlag = css.NewStyle().
		Bold(true).
		Padding(0, 2).
		Background(css.Color("#28282b")).
		Render("Dictionary")

	var txtInput = func() string {
		it.TextInput.Prompt = prefixStyle("Search: ")
		it.TextInput.TextStyle = css.NewStyle().Foreground(css.Color(yellow))

		return it.TextInput.View()
	}

	return box.Render(
		Div(
			titleApp+dictFlag,
			"",
			txtInput(),
			"",
			fontColor(it.Dictionary.Word, white),
			fontColor(it.Dictionary.PartOfSpeech, yellow),
			fontColor(it.Dictionary.Definition, blue),
			fontColor(it.Dictionary.Example, green),
			fontColor(it.Dictionary.Synonyms, purple),
			"",
			infoHelp,
		),
	)
}
