package tabs

import (
	"giobs/bs"
	"giobs/bs/label"
	"giobs/bs/theme"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Tab struct {
	btn     widget.Clickable
	content layout.Widget
	th      *theme.Theme

	title label.LabelStyle
}

func NewTab(th *theme.Theme, title string, content layout.Widget) *Tab {
	return &Tab{
		content: content,
		title:   bs.H6(th, title),
	}
}

func (t *Tab) LayoutTab(gtx layout.Context) layout.Dimensions {
	return layout.UniformInset(unit.Dp(8)).Layout(gtx,
		t.title.Layout,
	)
}

// nil pointer
// func (t *Tab) Layout(gtx layout.Context) layout.Dimensions {
// 	var tabWidth int

// 	return layout.Stack{Alignment: layout.S}.Layout(gtx,
// 		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
// 			dims := button.Clickable(gtx, &t.btn, func(gtx layout.Context) layout.Dimensions {
// 				return t.LayoutTab(gtx)
// 			})
// 			tabWidth = dims.Size.X
// 			return dims
// 		}),
// 		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
// 			// TODO
// 			// if t.selected != tabIdx {
// 			// 	return layout.Dimensions{}
// 			// }
// 			tabHeight := gtx.Dp(unit.Dp(4))
// 			tabRect := image.Rect(0, 0, tabWidth, tabHeight)
// 			paint.FillShape(gtx.Ops, t.th.Palette.ContrastBg, clip.Rect(tabRect).Op())
// 			return layout.Dimensions{
// 				Size: image.Point{X: tabWidth, Y: tabHeight},
// 			}
// 		}),
// 	)
// }

func (t *Tab) LayoutContent(gtx layout.Context) layout.Dimensions {
	return t.content(gtx)
}
