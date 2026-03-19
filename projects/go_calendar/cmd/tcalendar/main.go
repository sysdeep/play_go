package main

import (
	"fmt"
	"os"
	root "tcalendar"
	"tcalendar/internal/configuration"
	"tcalendar/internal/tui/application"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// // parse args
	// 	args := parseArgs()

	// 	if args.Version {
	// 		fmt.Println(sbadm.AppVersion)
	// 		os.Exit(0)
	// 	}

	conf := configuration.Configuration{
		AppVersion: root.AppVersion,
	}

	// TUI
	p := tea.NewProgram(application.New(conf))

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
