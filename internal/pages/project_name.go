package pages

import (
	"charm.land/bubbles/v2/textinput"
	"env-create/internal/services"
)

// 1st page
type ProjectNamePageMsg struct{}

func InitProjectNamePage(srv *services.ServicesStore) textinput.Model {
	defaultName := srv.ProjectName.GetProjectName()
	ti := textinput.New()
	ti.SetValue(defaultName)
	ti.Prompt = ""
	ti.CharLimit = 128
	ti.SetWidth(40)
	ti.Focus()

	return ti
}
