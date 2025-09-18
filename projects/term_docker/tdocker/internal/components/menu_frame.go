package components

import (
	"strings"
	"tdocker/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: сейчас выбор страницы и пункта меню никак не связаны друг с другом
type MenuFrame struct {
	options       []MenuFrameItem
	selectedIndex int
	isActive      bool
	maxWidth      int
}

type MenuFrameItem struct {
	Text    string
	OnPress func() tea.Msg
}

type toggleCasingMsg struct{}
type exitMsg struct{}

func NewMenuFrame(options []MenuFrameItem) MenuFrame {
	return MenuFrame{
		options:  options,
		maxWidth: 20,
	}
}

func (m MenuFrame) Init() tea.Cmd {
	// ????
	// return tea.SetWindowTitle("Grocery List")
	return nil
}

func (m MenuFrame) Update(msg tea.Msg) (MenuFrame, tea.Cmd) {

	// первая реакция на фокус
	switch msg := msg.(type) {

	case tui.FocusMsg:
		// fmt.Print(msg)
		m.isActive = msg.Focus == tui.FOCUS_MENU
		return m, nil
	}

	// если фокус потерян - не реагируем на события
	if !m.isActive {
		return m, nil
	}

	switch msg := msg.(type) {

	// case tui.FocusMsg:
	// 	// fmt.Print(msg)
	// 	m.isActive = msg.Focus == tui.FOCUS_MENU
	// 	return m, nil

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

func (m MenuFrame) View() string {

	wStyle := lipgloss.NewStyle().Width(m.maxWidth)

	var options []string
	for i, o := range m.options {
		preStr := "   "
		if i == m.selectedIndex {
			preStr = "-> "

		}
		options = append(options, wStyle.Render(preStr+o.Text))
	}

	// 	body := fmt.Sprintf(`%s

	// Press enter/return to select a list item, arrow keys to move, or Ctrl+C to exit.`,
	// 		strings.Join(options, "\n"))

	body := strings.Join(options, "\n")

	borderStyle := MakeFocusedBorder(m.isActive)

	return borderStyle.Render(body)
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
