package layout

import (
	"app/translate"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Model: variáveis persistentes durante a renderização
type LoroApp struct {
	textInput     textinput.Model
	textTranslate string
	switchLang    bool
	err           error
}

// SetLoroApp: inicializa o model com valores padrão
func SetLoroApp() LoroApp {
	var ti = textinput.New()
	ti.Placeholder = "..."
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 100

	return LoroApp{
		textInput:     ti,
		textTranslate: "...",
		switchLang:    false,
		err:           nil,
	}
}

func (it LoroApp) Init() tea.Cmd {
	return textinput.Blink
}

// LoroApp.Update() tea.Model, tea.Cmd: atualiza o model com base na mensagem recebida
func (it LoroApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

		case tea.KeyCtrlX:
			it.textInput.Reset()
		}

	case error:
		it.err = msg
		return it, nil
	}

	it.textInput, cmd = it.textInput.Update(msg)
	return it, cmd
}
