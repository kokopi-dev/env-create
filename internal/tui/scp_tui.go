package tui

import (
	"env-create/internal/services"

	"charm.land/bubbles/v2/textinput"
)

const (
	fieldUsername   = 0
	fieldHost       = 1
	fieldRemotePath = 2
	fieldCount      = 3
)

type ScpTUIInterface struct {
	Inputs       [fieldCount]textinput.Model
	ActiveField  int
	Accepted     bool
	Quitting     bool
	ShowResult   bool
	ResultOutput string
	ResultErr    error
	Services     *services.ServicesStore
	WindowWidth  int
	WindowHeight int
}

func NewScpTUIInterface(srv *services.ServicesStore) ScpTUIInterface {
	return ScpTUIInterface{
		Services: srv,
	}
}
