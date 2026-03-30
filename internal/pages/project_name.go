package pages

import (
	"env-create/internal/services"
	"env-create/internal/styles"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type ProjectNamePageMsg struct{}

type ProjectNameAcceptedMsg struct{ Value string }

type ProjectNameModel struct {
	input textinput.Model
}

func InitProjectNamePage(srv *services.ServicesStore) ProjectNameModel {
	ti := textinput.New()
	ti.SetValue(srv.ProjectName.GetProjectName())
	ti.Prompt = ""
	ti.CharLimit = 128
	ti.SetWidth(40)
	ti.Focus()

	return ProjectNameModel{input: ti}
}

func (p ProjectNameModel) Update(msg tea.Msg) (ProjectNameModel, tea.Cmd) {
	if msg, ok := msg.(tea.KeyPressMsg); ok {
		if msg.String() == "enter" {
			return p, func() tea.Msg { return ProjectNameAcceptedMsg{Value: p.input.Value()} }
		}
	}

	var cmd tea.Cmd
	p.input, cmd = p.input.Update(msg)
	return p, cmd
}

func (p ProjectNameModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Left,
		styles.InputLabelStyle.Render("Project name"),
		p.input.View(),
	)
}

func (p ProjectNameModel) Hints() string {
	return styles.FooterHint("enter", "confirm") +
		styles.FooterSep() +
		styles.FooterHint("esc", "cancel") +
		styles.FooterSep() +
		styles.FooterHint("ctrl+c", "quit")
}
