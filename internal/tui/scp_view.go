package tui

import (
	"env-create/internal/styles"
	"fmt"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)


var scpFieldLabels = [fieldCount]string{
	"Username",
	"Remote host",
	"Remote path",
}

func (m ScpTUIInterface) View() tea.View {
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

	var topContent string
	var footer string

	switch m.Page {
	case scpPageCachePrompt:
		topContent = styles.CardInnerStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				styles.CardTitleStyle.Render("✦  env-create"),
				styles.CardSubtitleStyle.Render("Send .env to remote server"),
				"",
				styles.InputLabelStyle.Render("No configs found."),
				styles.CardSubtitleStyle.Render("Save your configs for later?"),
				"",
				styles.HintStyle.Render("[y/n]"),
			),
		)
		footer = styles.FooterStyle.Render(
			footerHint("y", "save") +
				footerSep() +
				footerHint("n", "skip") +
				footerSep() +
				footerHint("esc", "cancel"),
		)

	case scpPageResult:
		var body string
		if m.ResultErr != nil {
			body = styles.ErrorStyle.Render("Transfer failed.") + "\n" + m.ResultErr.Error()
			if m.ResultOutput != "" {
				body += "\n" + m.ResultOutput
			}
		} else {
			out := m.ResultOutput
			if out == "" {
				out = "Done."
			}
			body = styles.ConfirmStyle.Render("Transfer complete.") + "\n" + out
		}

		topContent = styles.CardInnerStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				styles.CardTitleStyle.Render("✦  env-create"),
				styles.CardSubtitleStyle.Render("Send .env to remote server"),
				body,
			),
		)
		footer = styles.FooterStyle.Render(footerHint("enter", "close"))

	default: // scpPageForm
		var rows []string
		rows = append(rows,
			styles.CardTitleStyle.Render("✦  env-create"),
			styles.CardSubtitleStyle.Render("Send .env to remote server"),
		)
		for i := 0; i < fieldCount; i++ {
			labelStyle := styles.InputLabelStyle
			if i == m.ActiveField {
				labelStyle = styles.InputLabelActiveStyle
			}
			rows = append(rows, labelStyle.Render(scpFieldLabels[i]))
			rows = append(rows, m.Inputs[i].View())
			if i < fieldCount-1 {
				rows = append(rows, "")
			}
		}
		topContent = styles.CardInnerStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))

		stepHint := styles.FooterDescStyle.Render(fmt.Sprintf("step %d/%d", m.ActiveField+1, fieldCount))
		hints := footerHint("enter", "next") +
			footerSep() +
			footerHint("esc", "cancel") +
			footerSep() +
			stepHint
		footer = styles.FooterStyle.Render(hints)
	}

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
