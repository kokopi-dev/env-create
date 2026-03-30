package tui

import (
	"env-create/internal/styles"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type page interface {
	View() string
	Hints() string
}

func (m TUIInterface) activePage() page {
	switch m.Page {
	case "home":
		return m.Pages.Home
	case "project-name":
		return m.Pages.ProjectName
	}
	return m.Pages.Home
}

// main render function
// helps render the overall TUI
func (m TUIInterface) View() tea.View {
	if m.Quitting {
		return tea.NewView("")
	}

	w := m.WindowWidth
	h := m.WindowHeight
	if w == 0 {
		w = 80
	}
	if h == 0 {
		h = 24
	}

	p := m.activePage()

	topContent := styles.CardInnerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.CardTitleStyle.Render("✦  env-create"),
			styles.CardSubtitleStyle.Render("Set up your project environment"),
			p.View(),
		),
	)

	footer := styles.FooterStyle.Render(p.Hints())

	card := styles.CardStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			topContent,
			footer,
		),
	)

	cardHeight := lipgloss.Height(card)
	topPad := max((h-cardHeight)/2, 0)

	centeredCard := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		PaddingTop(topPad).
		Render(card)

	v := tea.NewView(centeredCard)
	v.AltScreen = true
	return v
}
