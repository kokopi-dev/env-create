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

type scpPage int

const (
	scpPageInit        scpPage = iota
	scpPageCachePrompt scpPage = iota
	scpPageForm        scpPage = iota
	scpPageResult      scpPage = iota
)

type ScpTUIInterface struct {
	Inputs       [fieldCount]textinput.Model
	ActiveField  int
	Accepted     bool
	Quitting     bool
	Page         scpPage
	ResultOutput string
	ResultErr    error
	SaveConfigs  bool
	Services     *services.ServicesStore
	WindowWidth  int
	WindowHeight int
}

func NewScpTUIInterface(srv *services.ServicesStore) ScpTUIInterface {
	return ScpTUIInterface{
		Services: srv,
	}
}
