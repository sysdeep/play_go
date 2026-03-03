package page

import (
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Router struct {
	pages   map[any]Page
	current any

	// drawer
	nav     component.NavDrawer
	navAnim component.VisibilityAnimation

	// app bar
	appBar *component.AppBar
}

func NewRouter() *Router {

	nav := component.NewNav("Title", "subtitle")

	na := component.VisibilityAnimation{
		State:    component.Visible,
		Duration: time.Millisecond * 250,
	}

	modal := component.NewModal()
	bar := component.NewAppBar(modal)
	// bar.NavigationIcon = icon.MenuIcon

	return &Router{
		pages:   make(map[any]Page),
		nav:     nav,
		navAnim: na,

		appBar: bar,
	}
}

func (r *Router) Register(tag any, page Page) {
	r.pages[tag] = page

	navItem := page.NavItem()
	navItem.Tag = tag // TODO: move to page
	if r.current == any(nil) {
		r.current = tag
		r.appBar.Title = navItem.Name
		// 	r.AppBar.SetActions(p.Actions(), p.Overflow())
	}
	r.nav.AddNavItem(navItem)
}

func (r *Router) SwitchTo(tag any) {
	p, ok := r.pages[tag]
	if !ok {
		return
	}
	navItem := p.NavItem()
	r.current = tag
	r.appBar.Title = navItem.Name
	// r.AppBar.SetActions(p.Actions(), p.Overflow())
}

func (r *Router) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	// for _, event := range r.AppBar.Events(gtx) {
	// 	switch event := event.(type) {
	// 	case component.AppBarNavigationClicked:
	// 		if r.NonModalDrawer {
	// 			r.NavAnim.ToggleVisibility(gtx.Now)
	// 		} else {
	// 			r.ModalNavDrawer.Appear(gtx.Now)
	// 			r.NavAnim.Disappear(gtx.Now)
	// 		}
	// 	case component.AppBarContextMenuDismissed:
	// 		log.Printf("Context menu dismissed: %v", event)
	// 	case component.AppBarOverflowActionClicked:
	// 		log.Printf("Overflow action selected: %v", event)
	// 	}
	// }

	if r.nav.NavDestinationChanged() {
		r.SwitchTo(r.nav.CurrentNavDestination())
	}

	paint.Fill(gtx.Ops, th.Palette.Bg)
	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,

			// main nav
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X /= 5
				return r.nav.Layout(gtx, th, &r.navAnim)
			}),

			// content
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return r.pages[r.current].Layout(gtx, th)
			}),
		)
	})

	bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return r.appBar.Layout(gtx, th, "Menu", "Actions")
	})

	flex := layout.Flex{Axis: layout.Vertical}

	// if r.BottomBar {
	// 	flex.Layout(gtx, content, bar)
	// } else {
	// 	flex.Layout(gtx, bar, content)
	// }

	flex.Layout(gtx, bar, content)

	// r.ModalLayer.Layout(gtx, th)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
