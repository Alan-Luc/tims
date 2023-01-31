package main

import (
	"fmt"
	"os"

	"github.com/Alan-Luc/tims/pkg/listing"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No search param was provided")
		os.Exit(1)
	}
	query := os.Args[1]
	if _, err := tea.NewProgram(listing.InitialModel(query)).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
