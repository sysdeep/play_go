package tabview

import (
	"giobs/bs/misc"
	"giobs/bs/theme"
	"image"
	"image/color"

	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

type TabItem struct {
	// Title of the tab.
	Title func(gtx C, th *theme.Theme) D
	// Main part of the tab content.
	Widget func(gtx C, th *theme.Theme) D
	// Title padding of the tab item.
	Inset     layout.Inset
	alignment layout.Direction
	click     gesture.Click
	btn       widget.Clickable
	hovering  bool
	selected  bool
}

func NewTabItem(inset layout.Inset, title, wgt func(gtx C, th *theme.Theme) D) *TabItem {
	return &TabItem{
		Title:  title,
		Widget: wgt,
		Inset:  inset,
	}
}

// func SimpleTabItem(inset layout.Inset, title string, wgt func(gtx C, th *theme.Theme) D) *TabItem {
// 	return &TabItem{
// 		Title: func(gtx C, th *theme.Theme) D {
// 			label := material.Label(th.Theme, th.TextSize, title)
// 			label.Font.Weight = font.Medium
// 			return label.Layout(gtx)
// 		},
// 		Widget: wgt,
// 		Inset:  inset,
// 	}
// }

func (item *TabItem) Update(gtx layout.Context) bool {
	for {
		event, ok := gtx.Event(
			pointer.Filter{Target: item, Kinds: pointer.Enter | pointer.Leave},
		)
		if !ok {
			break
		}

		switch event := event.(type) {
		case pointer.Event:
			switch event.Kind {
			case pointer.Enter:
				item.hovering = true
			case pointer.Leave:
				item.hovering = false
			case pointer.Cancel:
				item.hovering = false
			}
		}
	}

	var clicked bool
	for {
		e, ok := item.click.Update(gtx.Source)
		if !ok {
			break
		}
		if e.Kind == gesture.KindClick {
			clicked = true
			item.selected = true
		}
	}

	return clicked
}

func (item *TabItem) LayoutTab(gtx C, th *theme.Theme) D {
	item.Update(gtx)

	macro := op.Record(gtx.Ops)
	dims := item.layoutTab(gtx, th)
	call := macro.Stop()

	rect := clip.Rect(image.Rectangle{Max: dims.Size})
	defer rect.Push(gtx.Ops).Pop()

	item.click.Add(gtx.Ops)
	// register tag
	event.Op(gtx.Ops, item)
	call.Add(gtx.Ops)

	return dims
}

func (item *TabItem) layoutTab(gtx C, th *theme.Theme) D {
	return layout.Background{}.Layout(gtx,
		func(gtx C) D {
			return item.layoutTabBackground(gtx, th)
		},
		func(gtx C) D {
			return item.Inset.Layout(gtx, func(gtx C) D {
				return item.alignment.Layout(gtx, func(gtx C) D {
					return item.Title(gtx, th)
				})
			})
		},
	)
}

func (item *TabItem) layoutTabBackground(gtx C, th *theme.Theme) D {
	var fill color.NRGBA

	// TODO
	if item.hovering {
		// fill = misc.WithAlpha(th.Palette.Fg, th.HoverAlpha)
		fill = misc.WithAlpha(th.Palette.Fg, 100)
	} else if item.selected {
		// fill = misc.WithAlpha(th.Palette.Fg, th.SelectedAlpha)
		fill = misc.WithAlpha(th.Palette.Fg, 150)
	}

	rr := gtx.Dp(unit.Dp(4))
	rect := clip.RRect{
		Rect: image.Rectangle{
			Max: image.Point{X: gtx.Constraints.Min.X, Y: gtx.Constraints.Min.Y},
		},
		NE: rr,
		SE: rr,
		NW: rr,
		SW: rr,
	}
	paint.FillShape(gtx.Ops, fill, rect.Op(gtx.Ops))
	return layout.Dimensions{Size: gtx.Constraints.Min}
}

func (item *TabItem) LayoutWidget(gtx C, th *theme.Theme) D {
	return item.Widget(gtx, th)
}
