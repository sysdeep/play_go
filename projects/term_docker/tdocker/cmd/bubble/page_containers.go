package main

import tea "github.com/charmbracelet/bubbletea"

type pageContainers struct{}

func (p pageContainers) Init() tea.Cmd {
	return nil
}

func (p pageContainers) View() string {
	return "This is containers page"
}

func (p pageContainers) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}
