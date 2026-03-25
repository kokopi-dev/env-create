package tui

import (
	"env-create/internal/styles"
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m TUIInterface) View() tea.View {
	if m.Quitting {
		return tea.NewView("")
	}

	content := fmt.Sprintf(
		"%s %s\n\n%s\n",
		styles.LabelStyle.Render("Project Name:"),
		m.Input.View(),
		styles.HintStyle.Render("Edit the name, then press Enter to confirm • Esc to cancel"),
	)
	return tea.NewView(content)
}
