package tabs

import (
	"giobs/bs/button"
	"giobs/bs/theme"
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

/*
see examples
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

				var tabWidth int
				return layout.Stack{Alignment: layout.S}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := button.Clickable(gtx, &tb.btn, func(gtx layout.Context) layout.Dimensions {
							return t.tabs[tabIdx].LayoutTab(gtx)
						})
						tabWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						if t.selected != tabIdx {
							return layout.Dimensions{}
						}
						tabHeight := gtx.Dp(unit.Dp(4))
						tabRect := image.Rect(0, 0, tabWidth, tabHeight)
						paint.FillShape(gtx.Ops, t.th.Palette.ContrastBg, clip.Rect(tabRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: tabWidth, Y: tabHeight},
						}
					}),
				)
			})
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// return bs.H2(t.th, "aaaa").Layout(gtx)
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
