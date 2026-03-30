package pages

import "charm.land/bubbles/v2/textinput"

type ScpPageMsg struct{}

type ScpResultMsg struct {
	Output string
	Err    error
}

func InitUsernameInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "admin"
	ti.Prompt = ""
	ti.CharLimit = 256
	ti.SetWidth(40)
	return ti
}

func InitHostInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "192.168.1.1"
	ti.Prompt = ""
	ti.CharLimit = 256
	ti.SetWidth(40)
	return ti
}

func InitRemotePathInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "/home/admin/app"
	ti.Prompt = ""
	ti.CharLimit = 256
	ti.SetWidth(40)
	return ti
}
