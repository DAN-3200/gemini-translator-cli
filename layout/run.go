package layout

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	var program = tea.NewProgram(SetLoroApp())
	var _, err = program.Run()
	if err != nil { // tratamento erro
		fmt.Printf("Erro: %v/n", err)
	}
}
