package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	m := menu{
		options: []menuItem{
			menuItem{
				text:    "one",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			menuItem{
				text:    "two",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
			menuItem{
				text:    "three",
				onPress: func() tea.Msg { return toggleCasingMsg{} },
			},
		},
	}

	// p := tea.NewProgram(NewSimplePage("hello"))
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
