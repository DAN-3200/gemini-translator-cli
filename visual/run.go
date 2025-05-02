package layout

import (
	"fmt"

	"app/visual/layout"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	var program = tea.NewProgram(layout.SetApp())
	var _, err = program.Run()
	if err != nil {
		fmt.Printf("Erro: %v/n", err)
	}
}
