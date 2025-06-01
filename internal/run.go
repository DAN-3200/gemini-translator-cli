package internal

import (
	"app/internal/agentAI"
	"app/internal/engine"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	agentAI.InitGemini()

	var program = tea.NewProgram(engine.SetApp())
	var _, err = program.Run()
	if err != nil {
		fmt.Printf("Erro: %v/n", err)
	}
}
