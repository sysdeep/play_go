package tabs

import (
	"giobs/bs"
	"giobs/bs/button"
	"giobs/bs/label"
	"giobs/bs/theme"
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Tab struct {
	btn     widget.Clickable
	content layout.Widget
	th      *theme.Theme

	title    label.LabelStyle
	selected bool
}

func NewTab(th *theme.Theme, title string, content layout.Widget) *Tab {
	return &Tab{
		content: content,
		title:   bs.H6(th, title),
		th:      th,

		selected: false,
	}
}

func (t *Tab) layoutTab(gtx layout.Context) layout.Dimensions {
	// return layout.UniformInset(unit.Dp(8)).Layout(gtx,

	lay := layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(8), Right: unit.Dp(8)}

	return lay.Layout(gtx,
		t.title.Layout,
	)
}

func (t *Tab) Layout(gtx layout.Context) layout.Dimensions {
	// t.Update(gtx)

	var tabWidth int
	var tabHeight int
	return layout.Stack{Alignment: layout.S}.Layout(gtx,

		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			dims := button.Clickable(gtx, &t.btn, func(gtx layout.Context) layout.Dimensions {
				return t.layoutTab(gtx)
			})
			tabWidth = dims.Size.X
			tabHeight = dims.Size.Y
			return dims
		}),

		layout.Stacked(func(gtx layout.Context) layout.Dimensions {

			tabRect := image.Rect(0, 0, tabWidth, tabHeight)
			const r = 8
			rrect := clip.RRect{Rect: tabRect, NE: r, NW: r}
			paint.FillShape(gtx.Ops, color.NRGBA{A: 255, R: 0, G: 0, B: 0},
				clip.Stroke{
					Path:  makeTabPath(gtx.Ops, rrect, t.selected),
					Width: 1,
				}.Op(),
			)

			return layout.Dimensions{
				Size: image.Point{X: tabWidth, Y: tabHeight},
			}
		}),

		// layout.Stacked(func(gtx layout.Context) layout.Dimensions {
		// 	if !t.selected {
		// 		return layout.Dimensions{}
		// 	}
		// 	tabHeight := gtx.Dp(unit.Dp(4))
		// 	tabRect := image.Rect(0, 0, tabWidth, tabHeight)

		// 	paint.FillShape(gtx.Ops, t.th.Palette.ContrastBg, clip.Rect(tabRect).Op())

		// 	return layout.Dimensions{
		// 		Size: image.Point{X: tabWidth, Y: tabHeight},
		// 	}
		// }),
	)
}

func (t *Tab) LayoutContent(gtx layout.Context) layout.Dimensions {
	return t.content(gtx)
}

// func (t *Tab) Update(gtx layout.Context) {
// 	t.selected = t.btn.Clicked(gtx)
// }

// func Clickable(gtx layout.Context, button *widget.Clickable, w layout.Widget) layout.Dimensions {
// 	return button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
// 		semantic.Button.Add(gtx.Ops)
// 		return layout.Background{}.Layout(gtx,
// 			func(gtx layout.Context) layout.Dimensions {
// 				defer clip.Rect{Max: gtx.Constraints.Min}.Push(gtx.Ops).Pop()
// 				if button.Hovered() || gtx.Focused(button) {
// 					paint.Fill(gtx.Ops, f32color.Hovered(color.NRGBA{}))
// 				}
// 				for _, c := range button.History() {
// 					drawInk(gtx, c)
// 				}
// 				return layout.Dimensions{Size: gtx.Constraints.Min}
// 			},
// 			w,
// 		)
// 	})
// }

// Path returns the PathSpec for the top rounded .
func makeTabPath(ops *op.Ops, rr clip.RRect, active bool) clip.PathSpec {
	var p clip.Path
	p.Begin(ops)

	// https://pomax.github.io/bezierinfo/#circles_cubic.
	const q = 4 * (math.Sqrt2 - 1) / 3
	const iq = 1 - q

	// se, sw, nw, ne := float32(rr.SE), float32(rr.SW), float32(rr.NW), float32(rr.NE)
	_, _, nw, ne := float32(rr.SE), float32(rr.SW), float32(rr.NW), float32(rr.NE)
	w, n, e, s := float32(rr.Rect.Min.X), float32(rr.Rect.Min.Y), float32(rr.Rect.Max.X), float32(rr.Rect.Max.Y)

	p.MoveTo(f32.Point{X: w, Y: s})

	p.LineTo(f32.Point{X: w, Y: n + nw}) // to N
	p.CubeTo(                            // NW
		f32.Point{X: w, Y: n + nw*iq},
		f32.Point{X: w + nw*iq, Y: n},
		f32.Point{X: w + nw, Y: n})

	p.LineTo(f32.Point{X: e - ne, Y: n}) // to E
	p.CubeTo(                            // NE
		f32.Point{X: e - ne*iq, Y: n},
		f32.Point{X: e, Y: n + ne*iq},
		f32.Point{X: e, Y: n + ne})

	p.LineTo(f32.Point{X: e, Y: s}) // to S

	if !active {
		p.LineTo(f32.Point{X: w, Y: s})
	}
	return p.End()

}
