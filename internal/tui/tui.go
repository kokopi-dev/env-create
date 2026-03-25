package tui

import (
	"env-create/internal/services"

	"charm.land/bubbles/v2/textinput"
)
type TUIInterface struct {
	Input    textinput.Model
	Quitting bool
	Accepted bool
	Services *services.ServicesStore
}

func NewTUIInterface(servicesStore *services.ServicesStore) TUIInterface {
	return TUIInterface{Services: servicesStore}
}
