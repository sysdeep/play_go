package main

import (
	"tdocker/internal/components"
	"tdocker/internal/tui"
	"tdocker/internal/tui/master"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/client"
)

func main() {

	d_client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer d_client.Close()

	m := makeMainMenu()

	p := tea.NewProgram(
		// newSimplePage("This app is under construction"),
		// m,
		master.NewMasterModel(m, d_client),
	)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func makeMainMenu() components.MenuFrame {
	options := []components.MenuFrameItem{
		{
			Text:    "Containers",
			OnPress: tui.MakeGotoPageMsg(tui.PAGE_CONTAINERS),
		},
		{
			Text:    "Images",
			OnPress: tui.MakeGotoPageMsg(tui.PAGE_IMAGES),
		},
		{
			Text:    "Volumes",
			OnPress: tui.MakeGotoPageMsg(tui.PAGE_VOLUMES),
		},
		{
			Text:    "Networks",
			OnPress: tui.MakeGotoPageMsg(tui.PAGE_NETWORKS),
		},
	}

	return components.NewMenuFrame(options)

}

func makeDocker() {
	// test docker

}
