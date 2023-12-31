package portal

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	items := []list.Item{
		listItem{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		listItem{title: "Nutella", desc: "It's good on toast"},
		listItem{title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Fave Things"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}
