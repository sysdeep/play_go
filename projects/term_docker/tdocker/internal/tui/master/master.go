package master

import (
	"slices"
	"tdocker/internal/components"
	"tdocker/internal/tui"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/docker/client"
)

const gap = "\n\n"

type (
	errMsg error
)

type MasterModel struct {
	// viewport viewport.Model
	menu        components.MenuFrame
	senderStyle lipgloss.Style
	err         error
	sub         chan struct{}

	// pages
	pageGeometry *components.PageGeometry
	npages       map[int]tea.Model
	page         int

	// focus
	focuses []string
}

// https://github.com/charmbracelet/bubbletea/tree/main/examples/chat
func NewMasterModel(menu components.MenuFrame, dockerCli *client.Client) MasterModel {

	pageGeometry := &components.PageGeometry{
		MaxWidth:  10,
		MaxHeight: 10,
	}

	return MasterModel{
		menu:        menu,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
		sub:         make(chan struct{}),

		npages: map[int]tea.Model{
			tui.PAGE_CONTAINERS: components.NewPageContainers(dockerCli, pageGeometry),
			tui.PAGE_IMAGES:     components.NewPageImages(dockerCli, pageGeometry),
			tui.PAGE_VOLUMES:    components.NewPageVolumes(pageGeometry),
			tui.PAGE_NETWORKS:   components.NewPageNetworks(pageGeometry),
		},
		page:         tui.PAGE_CONTAINERS,
		pageGeometry: pageGeometry,

		focuses: []string{
			tui.FOCUS_MENU,
			tui.FOCUS_PAGE,
		},
	}
}

func (m MasterModel) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		tui.MakeFocusMsg(tui.FOCUS_MENU), // setup current focus
		listenForActivity(m.sub), waitForActivity(m.sub),

		// шлём событие что надо обновить страничку
		func() tea.Msg { return tui.NeedRefreshMsg{Page: tui.PAGE_CONTAINERS} },
	)

}

func (m MasterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = []tea.Cmd{}

	// menu
	menu, mCmd := m.menu.Update(msg)
	cmds = append(cmds, mCmd)
	m.menu = menu

	// current page
	p, pCmp := m.npages[m.page].Update(msg)
	cmds = append(cmds, pCmp)
	m.npages[m.page] = p

	switch msg := msg.(type) {

	case PollMsg:
		// am.responses++                     // record external activity
		// m.npages[m.page]
		// fmt.Println("poll///")
		return m, waitForActivity(m.sub) // wait for next event

	// размеры
	case tea.WindowSizeMsg:

		m.pageGeometry.MaxWidth = msg.Width - 24
		m.pageGeometry.MaxHeight = msg.Height - 12
		return m, nil

		// switch pp := m.npages[m.page].(type) {
		// case components.PageContainers:
		// 	pp.MaxWidth = msg.Width - 24
		// 	pp.MaxHeight = msg.Height - 12
		// 	m.npages[m.page] = pp
		// 	return m, nil
		// }

	// m.viewport.Width = msg.Width
	// m.textarea.SetWidth(msg.Width)
	// m.viewport.Height = msg.Height - m.textarea.Height() - lipgloss.Height(gap)
	// m.viewport.Height = msg.Height - m.textarea.Height() - (lipgloss.Height(gap) * 2) - 10

	// if len(m.messages) > 0 {
	// 	// Wrap content before setting it.
	// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
	// }
	// m.viewport.GotoBottom()

	// goto page event handling
	case tui.GotoPageMsg:
		m.page = msg.Page

		return m, func() tea.Msg { return tui.NeedRefreshMsg{Page: msg.Page} }

	// keys handling
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		// case tea.KeyCtrlN:

		case tea.KeyTab:
			m.changeFocus()
			return m, tui.MakeFocusMsg(m.focuses[0])

			// case tea.KeyEnter:
			// 	m.messages = append(m.messages, m.senderStyle.Render("You: ")+m.textarea.Value())
			// 	m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			// 	m.textarea.Reset()
			// 	m.viewport.GotoBottom()
		}

		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(cmds...)
}

func (m *MasterModel) changeFocus() {
	rotate(m.focuses, 1)
}

// function to rotate array by k elems (3 reverse method)
func rotate(arr []string, k int) {
	slices.Reverse(arr[:k])
	slices.Reverse(arr[k:])
	slices.Reverse(arr)
}

var (
	// focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

func (m MasterModel) View() string {

	p_view := ""
	cp, ok := m.npages[m.page]
	if ok {
		p_view = cp.View()
	}

	m_style := blurredStyle
	// if m.focus == FOCUS_MENU {
	// 	m_style = focusedStyle
	// }

	p_style := blurredStyle
	// if m.focus == FOCUS_PAGE {
	// 	p_style = focusedStyle
	// }

	// test columns
	// main_view := lipgloss.JoinHorizontal(lipgloss.Top, m.menu.View(), m.pages[m.pageIndex].View())
	main_view := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m_style.Render(m.menu.View()),
		p_style.Render(p_view),
	)

	screen := lipgloss.JoinVertical(
		lipgloss.Top,
		main_view,
		gap,
		"Press Ctrl+C or Esc or q to exit",
	)

	return screen

	// попытки изобразить модальчик - неудачно.. все ждут 2 версию lipgloss где такой функционал обещали
	// over := lipgloss.NewStyle().
	// 	// Width(10).
	// 	// Height(10).
	// 	Padding(1, 2).
	// 	BorderStyle(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("202")).Render("Over text")

	// return screen + lipgloss.Place(100, 100, lipgloss.Center, lipgloss.Center, over)
	// return "s\no\nm\ne\n text" + lipgloss.Place(10, 10, lipgloss.Center, lipgloss.Center, over)
	// return lipgloss.Place(100, 10, lipgloss.Center, lipgloss.Center, over)

	// return lipgloss.JoinVertical(lipgloss.Center, screen, over)

	// return fmt.Sprintf(
	// 	"%s%s%s",
	// 	// m.viewport.View(),
	// 	// "\nPress Ctrl+N to change page\n\n",
	// 	// m.menu.View(),
	// 	main_view,
	// 	gap,
	// 	gap,
	// 	// m.textarea.View(),
	// )
}
