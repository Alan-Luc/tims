package listing

import (
	"github.com/Alan-Luc/tims/pkg/img"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var imgs = Imgs

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(img.Img); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.scale):
				index := m.Index()
				imgs.Scale(index)
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.scale.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("scaled " + title + " by 2x"))
			case key.Matches(msg, keys.convertJPG):
				index := m.Index()
				imgs.Convert(index, "jpg")
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.convertJPG.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("converted " + title + " to jpg"))
			case key.Matches(msg, keys.convertPNG):
				index := m.Index()
				imgs.Convert(index, "png")
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.convertPNG.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("converted " + title + " to png"))
			case key.Matches(msg, keys.rotateCW):
				index := m.Index()
				imgs.Rotate(index, "90")
				if len(m.Items()) == 0 {
					keys.convertPNG.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("rotated " + title + " 90° clockwise"))
			case key.Matches(msg, keys.rotateCCW):
				index := m.Index()
				imgs.Rotate(index, "-90")
				if len(m.Items()) == 0 {
					keys.convertPNG.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("rotated " + title + " 90° counter clockwise"))
			case key.Matches(msg, keys.monochrome):
				index := m.Index()
				imgs.Monochrome(index)
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.convertPNG.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("converted " + title + " to B&W"))
			}
		}
		return nil
	}

	help := []key.Binding{
		keys.scale,
		keys.convertJPG,
		keys.convertPNG,
		keys.rotateCW,
		keys.rotateCCW,
		keys.monochrome,
	}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	scale      key.Binding
	convertJPG key.Binding
	convertPNG key.Binding
	rotateCW   key.Binding
	rotateCCW  key.Binding
	monochrome key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.scale,
		d.convertJPG,
		d.convertPNG,
		d.rotateCW,
		d.rotateCCW,
		d.monochrome,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.scale,
			d.convertJPG,
			d.convertPNG,
			d.rotateCW,
			d.rotateCCW,
			d.monochrome,
		},
	}
}

func NewDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		scale: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "scale"),
		),
		convertJPG: key.NewBinding(
			key.WithKeys("J"),
			key.WithHelp("J", "convert to JPG"),
		),
		convertPNG: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "convert to PNG"),
		),
		rotateCW: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "rotate image 90° clockwise"),
		),
		rotateCCW: key.NewBinding(
			key.WithKeys("R"),
			key.WithHelp("R", "rotate image 90° counter clockwise"),
		),
		monochrome: key.NewBinding(
			key.WithKeys("m"),
			key.WithHelp("m", "transform image to black and white"),
		),
	}
}
