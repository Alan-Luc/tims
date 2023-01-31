package listing

import (
	"fmt"
	"time"

	"github.com/Alan-Luc/tims/pkg/img"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// bubbles styling
var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("d7d7d7")).
			Background(lipgloss.Color("#4e5b9e")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#4e5b9e", Dark: "#848a84"}).
				Render
	Imgs = &img.Images{}
)

type listKeyMap struct {
	scale      key.Binding
	convertJpg key.Binding
	convertPng key.Binding
}

type Model struct {
	list         list.Model
	imgs         img.Images
	delegateKeys *delegateKeyMap
}

func InitialModel(query string) Model {
	delegateKeys := NewDelegateKeyMap()

	Imgs.Find(query)
	images := make([]list.Item, 0)
	for _, elem := range *Imgs {
		images = append(images, elem)
	}

	delegate := newItemDelegate(delegateKeys)
	imagesList := list.New(images, delegate, 0, 0)

	imagesList.Title = fmt.Sprintf("search results for: %s", query)
	imagesList.Styles.Title = titleStyle
	imagesList.StatusMessageLifetime = 10 * time.Second
	return Model{
		list:         imagesList,
		delegateKeys: delegateKeys,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return appStyle.Render(m.list.View())
}
