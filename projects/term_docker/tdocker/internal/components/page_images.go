package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PageImages struct {
	selectedIndex int
	options       []string
}

func NewPageImages() PageImages {
	return PageImages{
		selectedIndex: 0,
		options:       []string{"astra", "ubuntu"},
	}
}

func (p PageImages) Init() tea.Cmd {
	return nil
}

func (p PageImages) View() string {
	var options []string
	for i, o := range p.options {
		if i == p.selectedIndex {
			options = append(options, fmt.Sprintf("-> %s", o))
		} else {
			options = append(options, fmt.Sprintf("   %s", o))
		}
	}

	body := strings.Join(options, "\n")

	style := lipgloss.NewStyle().
		// BorderStyle(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.RoundedBorder())
		// .
		// BorderBackground(lipgloss.Color("240"))
		// BorderForeground(lipgloss.Color("240"))

	return style.Render(body)

}

func (p PageImages) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// case toggleCasingMsg:
	// 	return m.toggleSelectedItemCase(), nil
	// case exitMsg:
	// 	return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		// case "ctrl+c":
		// 	return m, tea.Quit
		case "down", "right", "up", "left":
			// return m.moveCursor(msg.String()), nil
			p.moveCursor(msg.String())
			return p, nil
			// case "enter", "return":
			// 	return m, m.options[m.selectedIndex].onPress
		}
	}
	return p, nil
}

func (p *PageImages) moveCursor(msg string) {
	switch msg {
	case "up", "left":
		p.selectedIndex--
	case "down", "right":
		p.selectedIndex++
	default:
		// do nothing
	}

	optCount := len(p.options)
	p.selectedIndex = (p.selectedIndex + optCount) % optCount
	// return m
}
