package main

import (
	"env-create/internal/services"
	"env-create/internal/styles"
	"env-create/internal/tui"
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	srv := services.NewServicesStore()
	m := tui.NewTUIInterface(srv)
	p := tea.NewProgram(m)

	result, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	final, ok := result.(tui.TUIInterface)
	if !ok || !final.Accepted {
		fmt.Println("Cancelled.")
		os.Exit(0)
	}

	projectName := final.Input.Value()
	fmt.Println(styles.ConfirmStyle.Render(projectName))
}
