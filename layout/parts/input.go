package parts

import (
	"fmt"

	"app/translate"

	"app/pkg"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	css "github.com/charmbracelet/lipgloss"
)

type errMsg error

type TranslateField struct {
	textInput     textinput.Model
	textTranslate string
	switchLang    bool
	err           error
}

func Use_TranslateField() TranslateField {
	ti := textinput.New()
	ti.Placeholder = "Type..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	return TranslateField{
		textInput:     ti,
		textTranslate: "",
		err:           nil,
		switchLang:    false,
	}
}

func (it TranslateField) Init() tea.Cmd {
	return textinput.Blink
}

func (it TranslateField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return it, tea.Quit
		case tea.KeyCtrlD:
			it.switchLang = !it.switchLang
		case tea.KeyEnter:
			it.textTranslate = translate.UseGemini(it.textInput.Value(), it.switchLang)
		}

	// We handle errors just like any other message
	case errMsg:
		it.err = msg
		return it, nil
	}

	it.textInput, cmd = it.textInput.Update(msg)
	return it, cmd
}

// --------------------------------------------------------------------------------------

const (
	red    = "#f61363" // #f61363 vermelho
	green  = "#50a231" // #50a231 verde
	yellow = "#e6ba27" // #e6ba27 amarelo
	white  = "#ffffff"
)

func (it TranslateField) View() string {
	var layout = func() string {
		return css.NewStyle().Border(css.NormalBorder(), true, false).Width(40).Padding(2).
			Render("%s \n\n %s \n\n > %s \n\n %s")
	}

	var titleApp = css.NewStyle().Bold(true).Foreground(css.Color(red)).Blink(true).
		Render("| 🦜 Loro App |")

	var langSwap = func() string {
		var en = css.NewStyle().Bold(true).Foreground(css.Color(pkg.Ternary(it.switchLang, white, red))).
			Render("en")
		var pt = css.NewStyle().Bold(true).Foreground(css.Color(pkg.Ternary(it.switchLang, red, white))).
			Render("pt")
		return fmt.Sprintf(" %s", pkg.Ternary(it.switchLang, pt+" -> "+en, en+" -> "+pt))
	}

	var textTranslate = css.NewStyle().Foreground(css.Color(pkg.Ternary(it.switchLang, green, yellow))).
		Render(it.textTranslate)

	var infoHelp = css.NewStyle().Foreground(css.Color("#404040")).
		Render("[Esc] Exit [Ctrl+d] Switch language")

	var styled_TextInput = it.textInput
	styled_TextInput.TextStyle = css.
		NewStyle().
		Foreground(css.Color(pkg.Ternary(it.switchLang, yellow, green)))

	return fmt.Sprintf(
		layout(),
		titleApp+langSwap(),
		styled_TextInput.View(),
		textTranslate,
		infoHelp,
	)
}
