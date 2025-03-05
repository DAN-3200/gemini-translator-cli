package layout

import (
	"fmt"

	"app/layout/parts"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	var program = tea.NewProgram(parts.Use_TranslateField())
	var _, err = program.Run()
	if err != nil { // tratamento erro
		fmt.Printf("Erro: %v/n", err)
	}
}
