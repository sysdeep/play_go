package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuFrame struct {
	options       []MenuFrameItem
	selectedIndex int
	isActive      bool
}

type MenuFrameItem struct {
	Text    string
	OnPress func() tea.Msg
}

type toggleCasingMsg struct{}
type exitMsg struct{}

// type gotoPageMsg struct {
// 	Page int
// }

func NewMainFrame(options []MenuFrameItem) MenuFrame {
	return MenuFrame{
		options: options,
	}
}

func (m MenuFrame) Init() tea.Cmd {
	// ????
	return tea.SetWindowTitle("Grocery List")
}

func (m MenuFrame) View() string {
	var options []string
	for i, o := range m.options {
		if i == m.selectedIndex {
			options = append(options, fmt.Sprintf("-> %s", o.Text))
		} else {
			options = append(options, fmt.Sprintf("   %s", o.Text))
		}
	}

	// 	body := fmt.Sprintf(`%s

	// Press enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.`,
	// 		strings.Join(options, "\n"))

	body := strings.Join(options, "\n")

	color := "240"
	if m.isActive {
		color = "205"
	}

	style := lipgloss.NewStyle().
		// BorderStyle(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(color))

	return style.Render(body)
}

func (m MenuFrame) Update(msg tea.Msg) (MenuFrame, tea.Cmd) {
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
			return m, m.options[m.selectedIndex].OnPress
			// case "enter", "return":
			// 	return m, m.options[m.selectedIndex].onPress
		}
	}
	return m, nil
}

func (m *MenuFrame) moveCursor(msg string) {
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

func (m MenuFrame) toggleSelectedItemCase() MenuFrame {
	selectedText := m.options[m.selectedIndex].Text
	if selectedText == strings.ToUpper(selectedText) {
		m.options[m.selectedIndex].Text = strings.ToLower(selectedText)
	} else {
		m.options[m.selectedIndex].Text = strings.ToUpper(selectedText)
	}
	return m
}

func (m *MenuFrame) SetFocus(focus bool) {
	m.isActive = focus
}
