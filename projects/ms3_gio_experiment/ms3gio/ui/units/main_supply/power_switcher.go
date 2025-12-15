package mainsupply

import (
	"image"
	"image/color"
	"ms3gio/ui/palette"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
)

type powerSwitcher struct {
	state    bool
	hovered  bool
	pressed  bool
	bounding image.Rectangle
	pos      image.Point

	clickS  *widget.Clickable
	clicked bool
}

func newPowerSwitcher(pos image.Point) *powerSwitcher {
	return &powerSwitcher{
		bounding: image.Rect(0, 0, 24, 24),
		pos:      pos,
		clickS:   new(widget.Clickable),
	}
}

func (p *powerSwitcher) Layout(gtx layout.Context) layout.Dimensions {

	p.update(gtx)
	p.clickS.Update(gtx)
	defer op.Offset(p.pos).Push(gtx.Ops).Pop()

	body := clip.UniformRRect(p.bounding, 2).Op(gtx.Ops)

	// declare area for events
	area := clip.Rect(p.bounding).Push(gtx.Ops)
	event.Op(gtx.Ops, p)
	area.Pop()

	paint.FillShape(gtx.Ops, p.getColor(), body)

	return layout.Dimensions{Size: p.bounding.Max}
}

func (p *powerSwitcher) BoundingRect() image.Rectangle {
	return p.bounding
}

func (p *powerSwitcher) getColor() color.NRGBA {

	if p.pressed {
		return palette.Color(palette.Green, palette.P800)
	}

	if p.hovered {
		return palette.Color(palette.Green, palette.P400)
	}
	return palette.Color(palette.Green, palette.P600)

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

func (p *powerSwitcher) update(gtx layout.Context) {
	for {
		ev, ok := gtx.Source.Event(pointer.Filter{
			Target: p,
			Kinds:  pointer.Enter | pointer.Leave | pointer.Press | pointer.Release,
		})
		if !ok {
			break
		}

		if x, ok := ev.(pointer.Event); ok {
			switch x.Kind {
			case pointer.Enter:
				p.hovered = true
			case pointer.Leave:
				p.hovered = false
			case pointer.Press:
				p.pressed = true
				p.clicked = true
			case pointer.Release:
				p.pressed = false
			}
		}
	}
}

func (b *powerSwitcher) Clicked() bool {
	defer func() { b.clicked = false }()
	return b.clicked
}
