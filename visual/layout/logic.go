package layout

import (
	"app/engine"
	"app/pkg/utils"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// LoroApp.Init() tea.Cmd: Comandos Inicias da aplicação
func (it App) Init() tea.Cmd {
	return textinput.Blink
}

// LoroApp.Update() tea.Model, tea.Cmd: atualiza o model com base na mensagem recebida
func (it App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return it, tea.Quit
		case tea.KeyCtrlD:
			it.CtxTranslate.SwitchLang = !it.CtxTranslate.SwitchLang
		case tea.KeyEnter:
			if it.SwitchMode {
				it.CtxTranslate.Text = engine.UseTranslation(it.TextInput.Value(), it.CtxTranslate.SwitchLang)
			} else {
				it.CtxDict.Response = strings.TrimSpace(
					engine.UseDictionary(
						strings.TrimSpace(it.TextInput.Value()),
					),
				)
			}
		case tea.KeyCtrlX:
			it.TextInput.Reset()
		case tea.KeyCtrlA:
			it.SwitchMode = !it.SwitchMode
		}

	case error:
		it.Err = msg
		return it, nil
	}

	it.TextInput, cmd = it.TextInput.Update(msg)
	return it, cmd
}

// LoroApp.View() string: renderiza a interface a partir do dados model
func (it App) View() string {
	return utils.Ternary(it.SwitchMode, it.TranslatePage(), it.DictPage())
}
