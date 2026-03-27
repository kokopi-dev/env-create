package tui

import (
	"env-create/internal/pages"
	"env-create/internal/services"

	"charm.land/bubbles/v2/textinput"
)

type TUIInterface struct {
	Input        textinput.Model
	Home         pages.HomeModel
	Page         string
	Quitting     bool
	Accepted     bool
	Services     *services.ServicesStore
	WindowWidth  int
	WindowHeight int
}

func NewTUIInterface(servicesStore *services.ServicesStore) TUIInterface {
	return TUIInterface{Services: servicesStore}
}
