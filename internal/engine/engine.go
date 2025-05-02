package engine

import (
	"app/internal/agentAI"
	comp "app/internal/components"
	"app/pkg/utils"
	"fmt"

	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
)

func Run() {
	var program = tea.NewProgram(SetApp())
	var _, err = program.Run()
	if err != nil {
		fmt.Printf("Erro: %v/n", err)
	}
}

// -- App Methods

func (it App) Init() tea.Cmd {
	return textinput.Blink
}

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
				it.CtxTranslate.Text = agentAI.UseTranslation(it.TextInput.Value(), it.CtxTranslate.SwitchLang)
			} else {
				it.CtxDict.Response = strings.TrimSpace(
					agentAI.UseDictionary(
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

func (it App) View() string {
	return utils.Ternary(it.SwitchMode, comp.TranslatePage(it.CtxMain), comp.DictPage(it.CtxMain))
}
