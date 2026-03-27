package tui

import (
	"env-create/internal/pages"
	"env-create/internal/services"
)

type PageStore struct {
	Home        pages.HomeModel
	ProjectName pages.ProjectNameModel
}

type TUIInterface struct {
	Pages        PageStore
	Page         string
	Quitting     bool
	Accepted     bool
	AcceptedValue string
	Services     *services.ServicesStore
	WindowWidth  int
	WindowHeight int
}

// TUI entry point
// tui/update is the main render loop/flow
// tui/view is the main render function
// tui/init is the initial state (home page)
func NewTUIInterface(servicesStore *services.ServicesStore) TUIInterface {
	return TUIInterface{Services: servicesStore}
}
