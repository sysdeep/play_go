package application

import tea "github.com/charmbracelet/bubbletea"

const NO_PAGE = "---NOPAGE---"

type Router struct {
	pages map[string]tea.Model
	page  string
}

func newRouter() *Router {
	return &Router{
		pages: map[string]tea.Model{},
		page:  NO_PAGE,
	}
}

func (r *Router) Register(id string, page tea.Model) {
	r.pages[id] = page
	if r.page == NO_PAGE {
		r.page = id
	}
}

func (r *Router) GoTo(id string) {
	_, ok := r.pages[id]
	if ok {
		r.page = id
	}
}

func (r *Router) MustPage() tea.Model {
	return r.pages[r.page]
}
