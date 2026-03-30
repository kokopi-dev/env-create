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

	if len(os.Args) > 1 && os.Args[1] == "scp" {
		if !srv.Scp.EnvExists() {
			fmt.Fprintln(os.Stderr, "No .env file found in current directory.")
			os.Exit(1)
		}
		m := tui.NewScpTUIInterface(srv)
		p := tea.NewProgram(m)
		_, err := p.Run()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

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

	projectName := final.AcceptedValue
	fmt.Println(styles.ConfirmStyle.Render(projectName))
}
