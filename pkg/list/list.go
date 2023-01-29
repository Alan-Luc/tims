package list

import (
	"fmt"

	"github.com/Alan-Luc/tims/pkg/img"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  img.Images       // img.Images = []img
	cursor   int              // which items our cursor is pointing at
	selected map[int]struct{} // which items are selected
}

func InitialModel() model {
	images := &img.Images{}
	images.Grep("yotsu")
	fmt.Println(*images)

	return model{
		choices:  *images,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select an image to manipulate\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">"
		}

		// is this choice selected
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}
