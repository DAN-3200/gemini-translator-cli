package components

import (
	"app/internal/models"
	"strings"

	css "github.com/charmbracelet/lipgloss"
)

func DictPage(it models.CtxMain) string {
	var dictFlag = css.NewStyle().
		Bold(true).
		Padding(0, 2).
		Background(css.Color("#28282b")).
		Render("Dictionary")

	{ // TextInput
		it.TextInput.Prompt = prefixStyle("Search: ")
		it.TextInput.TextStyle = css.NewStyle().Foreground(css.Color(yellow))
	}

	content := css.JoinVertical(css.Left,
		titleApp()+dictFlag,
		"",
		it.TextInput.View(),
		"",
		fontColor(it.Dictionary.Word, white),
		fontColor(it.Dictionary.PartOfSpeech, yellow),
		fontColor(it.Dictionary.Definition, blue),
		fontColor(it.Dictionary.Example, green),
		fontColor(it.Dictionary.Synonyms, purple),
		"",
	)

	rows := strings.Count(content, "\n") + 2

	return box.Render(
		content,
		css.Place(it.Size.Width, it.Size.Height-rows, css.Left, css.Bottom, infoHelp),
	)
}
