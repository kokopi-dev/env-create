package tui

import (
	"env-create/internal/styles"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func footerHint(key, desc string) string {
	return styles.FooterKeyStyle.Render(key) +
		" " +
		styles.FooterDescStyle.Render(desc)
}

func footerSep() string {
	return styles.FooterSepStyle.Render(" · ")
}

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

	// top content
	topContent := styles.CardInnerStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.CardTitleStyle.Render("✦  env-create"),
			styles.CardSubtitleStyle.Render("Set up your project environment"),
			styles.InputLabelStyle.Render("Project name"),
			m.Input.View(),
		),
	)

	// footer content
	hints := footerHint("enter", "confirm") +
		footerSep() +
		footerHint("esc", "cancel") +
		footerSep() +
		footerHint("ctrl+c", "quit")
	footer := styles.FooterStyle.Render(hints)

	// entire card
	card := styles.CardStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			topContent,
			footer,
		),
	)

	// card padding and alignment
	cardHeight := lipgloss.Height(card)

	topPad := max((h - cardHeight) / 2, 0)

	centeredCard := lipgloss.NewStyle().
		Width(w).
		Align(lipgloss.Center).
		PaddingTop(topPad).
		Render(card)

	v := tea.NewView(centeredCard)
	v.AltScreen = true // overtake the entire terminal
	return v
}
