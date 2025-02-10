package components

import tea "github.com/charmbracelet/bubbletea"

type PageVolumes struct{}

func NewPageVolumes() PageVolumes {
	return PageVolumes{}
}

func (p PageVolumes) Init() tea.Cmd {
	return nil
}

func (p PageVolumes) View() string {
	return "This is volumes page"
}

func (p PageVolumes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return p, nil
}
