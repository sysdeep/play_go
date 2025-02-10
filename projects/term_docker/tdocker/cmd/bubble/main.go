package main

import (
	"tdocker/internal/components"
	"tdocker/internal/state"
	"tdocker/internal/tui/master"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	m := makeMainMenu()

	p := tea.NewProgram(
		// newSimplePage("This app is under construction"),
		// m,
		master.NewMasterModel(m),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func makeMainMenu() components.MenuFrame {
	options := []components.MenuFrameItem{
		{
			Text: "Containers",
			// onPress: func() tea.Msg { return toggleCasingMsg{} },
			OnPress: func() tea.Msg { return master.GotoPageMsg{state.PAGE_CONTAINERS} },
		},
		{
			Text: "Images",
			// OnPress: func() tea.Msg { return toggleCasingMsg{} },
			OnPress: func() tea.Msg { return master.GotoPageMsg{state.PAGE_IMAGES} },
		},
		{
			Text: "Volumes",
			// OnPress: func() tea.Msg { return toggleCasingMsg{} },
			OnPress: func() tea.Msg { return master.GotoPageMsg{state.PAGE_VOLUMES} },
		},
		{
			Text: "Networks",
			// OnPress: func() tea.Msg { return toggleCasingMsg{} },
			OnPress: func() tea.Msg { return master.GotoPageMsg{state.PAGE_NETWORKS} },
		},
		// {
		// 	text:    "exit",
		// 	onPress: func() tea.Msg { return exitMsg{} },
		// },
	}

	return components.NewMainFrame(options)

}

// new
