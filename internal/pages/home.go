package pages

import (
	"env-create/internal/styles"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type HomePageMsg struct{}

type HomeModel struct {
	options []string
	cursor  int
}

func InitHomePage() HomeModel {
	return HomeModel{
		options: []string{"create", "send", "get", "help"},
		cursor:  0,
	}
}

func (h HomeModel) Update(msg tea.Msg) (HomeModel, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		switch msg.String() {
		case "up":
			if h.cursor > 0 {
				h.cursor--
			}
		case "down":
			if h.cursor < len(h.options)-1 {
				h.cursor++
			}
		case "enter":
			if h.cursor == 0 {
				return h, func() tea.Msg { return ProjectNamePageMsg{} }
			}
			return h, tea.Quit
		}
	}
	return h, nil
}

func (h HomeModel) View() string {
	var items []string
	for i, opt := range h.options {
		if i == h.cursor {
			items = append(items, styles.ConfirmStyle.Render("> "+opt))
		} else {
			line := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("  "+opt)
			items = append(items, line)
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
