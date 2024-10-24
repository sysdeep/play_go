package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) View() string {
	var options []string
	for i, o := range m.options {
		if m.selectedIndex == i {
			options = append(options, fmt.Sprintf("-> %s", o.text))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.text))
		}
	}

	options_text := strings.Join(options, "\n")

	tpl := `%s

  Press enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.`
	return fmt.Sprintf(tpl, options_text)
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case toggleCasingMsg:
		return m.toggleMenuItem(), nil
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down", "up", "right", "left":
			return m.moveCursor(msg.(tea.KeyMsg)), nil
		case "enter", "return":
			return m, m.options[m.selectedIndex].onPress
		}
	}
	return m, nil
}

func (m menu) moveCursor(msg tea.KeyMsg) menu {
	switch msg.String() {
	case "up", "left":
		m.selectedIndex--
	case "down", "right":
		m.selectedIndex++
	default:
		// pass
	}

	optCount := len(m.options)
	m.selectedIndex = (m.selectedIndex + optCount) % optCount
	return m
}

func (m menu) toggleMenuItem() tea.Model {
	selectedText := m.options[m.selectedIndex].text
	if selectedText == strings.ToUpper(selectedText) {
		m.options[m.selectedIndex].text = strings.ToLower(selectedText)
	} else {
		m.options[m.selectedIndex].text = strings.ToUpper(selectedText)
	}
	return m
}
