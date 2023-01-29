package main

import (
	"fmt"
	"os"

	"github.com/Alan-Luc/tims/pkg/list"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(list.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("there has been an error: %v", err)
		os.Exit(1)
	}
}
