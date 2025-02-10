package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menu struct {
	options       []menuItem
	selectedIndex int
}

type menuItem struct {
	text    string
	onPress func() tea.Msg
}

type toggleCasingMsg struct{}
type exitMsg struct{}

type gotoPageMsg struct {
	Page int
}

func (m menu) Init() tea.Cmd {
	// ????
	return tea.SetWindowTitle("Grocery List")
}

func (m menu) View() string {
	var options []string
	for i, o := range m.options {
		if i == m.selectedIndex {
			options = append(options, fmt.Sprintf("-> %s", o.text))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.text))
		}
	}

	// 	body := fmt.Sprintf(`%s

	// Press enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.`,
	// 		strings.Join(options, "\n"))

	body := strings.Join(options, "\n")

	style := lipgloss.NewStyle().
		// BorderStyle(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	return style.Render(body)
}

func (m menu) Update(msg tea.Msg) (menu, tea.Cmd) {
	switch msg := msg.(type) {
	case toggleCasingMsg:
		return m.toggleSelectedItemCase(), nil
	case exitMsg:
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down", "right", "up", "left":
			// return m.moveCursor(msg.String()), nil
			m.moveCursor(msg.String())
			return m, m.options[m.selectedIndex].onPress
			// case "enter", "return":
			// 	return m, m.options[m.selectedIndex].onPress
		}
	}
	return m, nil
}

func (m *menu) moveCursor(msg string) {
	switch msg {
	case "up", "left":
		m.selectedIndex--
	case "down", "right":
		m.selectedIndex++
	default:
		// do nothing
	}

	optCount := len(m.options)
	m.selectedIndex = (m.selectedIndex + optCount) % optCount
	// return m
}

func (m menu) toggleSelectedItemCase() menu {
	selectedText := m.options[m.selectedIndex].text
	if selectedText == strings.ToUpper(selectedText) {
		m.options[m.selectedIndex].text = strings.ToLower(selectedText)
	} else {
		m.options[m.selectedIndex].text = strings.ToUpper(selectedText)
	}
	return m
}
