package main

import tea "github.com/charmbracelet/bubbletea"

type pageImages struct{}

func (p pageImages) Init() tea.Cmd {
	return nil
}

func (p pageImages) View() string {
	return "This is images page"
}

func (p pageImages) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}
