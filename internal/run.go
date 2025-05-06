package internal

import (
	"app/internal/engine"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	var program = tea.NewProgram(engine.SetApp())
	var _, err = program.Run()
	if err != nil {
		fmt.Printf("Erro: %v/n", err)
	}
}
