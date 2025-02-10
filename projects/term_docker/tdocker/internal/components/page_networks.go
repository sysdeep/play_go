package components

import tea "github.com/charmbracelet/bubbletea"

type PageNetworks struct{}

func NewPageNetworks() PageNetworks {
	return PageNetworks{}
}

func (p PageNetworks) Init() tea.Cmd {
	return nil
}

func (p PageNetworks) View() string {
	return "This is networks page"
}

func (p PageNetworks) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}
