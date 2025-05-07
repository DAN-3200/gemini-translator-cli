package engine

import (
	"app/internal/agentAI"
	comp "app/internal/components"
	u "app/pkg/utils"

	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/textinput"
)

func (it App) Init() tea.Cmd {
	return textinput.Blink
}

func (it App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		it.Size.Height = msg.Height
		it.Size.Width = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return it, tea.Quit
		case tea.KeyCtrlT:
			if it.SwitchMode {
				it.CtxTranslate.SwitchLang = !it.CtxTranslate.SwitchLang
			}
		case tea.KeyEnter:
			if it.SwitchMode {
				it.CtxTranslate.Text = agentAI.UseTranslation(it.TextInput.Value(), it.CtxTranslate.SwitchLang)
			} else {
				it.CtxDict.Dictionary = agentAI.UseDictionary(
					strings.TrimSpace(it.TextInput.Value()),
				)
			}
		case tea.KeyCtrlU:
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
	return u.Ternary(
		it.SwitchMode, comp.TranslatePage(it.CtxMain), comp.DictPage(it.CtxMain),
	)
}
