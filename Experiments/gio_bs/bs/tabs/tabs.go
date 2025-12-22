package tabs

import (
	"giobs/bs/theme"

	"gioui.org/layout"
	"gioui.org/unit"
)

/*

see examples

- https://getbootstrap.com/docs/5.3/components/navs-tabs/
*/

type Tabs struct {
	list     layout.List
	tabs     []*Tab
	selected int
	th       *theme.Theme
}

func NewTabs(th *theme.Theme, tabs ...*Tab) *Tabs {
	return &Tabs{
		th:       th,
		selected: 0,
		tabs:     tabs,
	}
}

func (t *Tabs) Layout(gtx layout.Context) layout.Dimensions {
	var flex = layout.Flex{Axis: layout.Vertical}
	return flex.Layout(gtx,

		// tabs list
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.list.Layout(gtx, len(t.tabs), func(gtx layout.Context, tabIdx int) layout.Dimensions {
				tb := t.tabs[tabIdx]

				if tb.btn.Clicked(gtx) {
					// if t.selected < tabIdx {
					// 	slider.PushLeft()
					// } else if tabs.selected > tabIdx {
					// 	slider.PushRight()
					// }
					t.selected = tabIdx

				}

				for i, ttab := range t.tabs {
					ttab.selected = i == t.selected
				}

				return tb.Layout(gtx)
			})
		}),

		// tab content
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.tabs[t.selected].LayoutContent(gtx)
		}),

		layout.Rigid(
			// The height of the spacer is 25 Device independent pixels
			layout.Spacer{Height: unit.Dp(25)}.Layout,
		),

		// layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {

		// 	return slider.Layout(gtx, func(gtx C) D {
		// 		fill(gtx, dynamicColor(tabs.selected), dynamicColor(tabs.selected+1))
		// 		return layout.Center.Layout(gtx,
		// 			material.H1(th, fmt.Sprintf("Tab content #%d", tabs.selected+1)).Layout,
		// 		)
		// 	})
		// }),
	)
}
