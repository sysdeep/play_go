package mainsupply

import (
	"image"
	"image/color"
	"ms3gio/ui/palette"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type body struct {
	// min      image.Point
	// max      image.Point
	hovered  bool
	bounding image.Rectangle
}

func newBody() *body {
	return &body{
		bounding: image.Rect(0, 0, 48, 64),
	}
}

func (b *body) Layout(gtx layout.Context) layout.Dimensions {

	b.update(gtx)
	// defer op.Offset(s.pos).Push(gtx.Ops).Pop()

	body := clip.UniformRRect(
		b.bounding,
		4,
	).Op(gtx.Ops)

	paint.FillShape(gtx.Ops, b.getColor(), body)
	// event.Op(gtx.Ops, b)

	area := clip.Rect(b.bounding).Push(gtx.Ops)
	event.Op(gtx.Ops, b)
	area.Pop()

	return layout.Dimensions{Size: b.bounding.Max}
}

func (b *body) BoundingRect() image.Rectangle {
	return b.bounding
}

func (b *body) getColor() color.NRGBA {
	if b.hovered {
		return palette.Color(palette.Amber, palette.P600)
	}

	return palette.Color(palette.Amber, palette.P800)

	// if d.model.IsBlock {
	// 	return color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	// }

	// if d.model.IsError {
	// 	return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	// }

	// if d.model.IsState {
	// 	return color.NRGBA{R: 0, G: 255, B: 40, A: 255}
	// }

	// return color.NRGBA{R: 0, G: 100, B: 40, A: 255}
}

func (b *body) update(gtx layout.Context) {
	for {
		ev, ok := gtx.Source.Event(pointer.Filter{
			Target: b,
			Kinds:  pointer.Press | pointer.Release,
		})

		if !ok {
			break
		}

		if x, ok := ev.(pointer.Event); ok {
			switch x.Kind {
			case pointer.Press:
				b.hovered = true
			case pointer.Release:
				b.hovered = false
			}
		}
	}
}
