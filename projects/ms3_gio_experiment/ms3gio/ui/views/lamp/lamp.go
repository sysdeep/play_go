package lamp

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type LampColor struct {
	Normal color.NRGBA
	Active color.NRGBA
}

type Lamp struct {
	rect image.Rectangle

	colorBorder LampColor
	colorBody   LampColor
	state       bool
}

func New(size int, colorBody, colorBorder LampColor) *Lamp {

	return &Lamp{
		rect:        image.Rect(0, 0, size, size),
		colorBorder: colorBorder,
		colorBody:   colorBody,
		state:       false,
	}
}

func (l *Lamp) Layout(gtx layout.Context) layout.Dimensions {
	// pos
	// defer op.Offset(s.pos).Push(gtx.Ops).Pop()

	// background
	circle := clip.Ellipse{
		Min: image.Pt(0, 0),
		Max: l.rect.Max,
	}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, l.getBorderColor(), circle)

	// foreground
	fc := clip.Ellipse{
		Min: image.Pt(2, 2),
		Max: l.rect.Max.Sub(image.Pt(2, 2)),
	}.Op(gtx.Ops)
	paint.FillShape(gtx.Ops, l.getBodyColor(), fc)

	return layout.Dimensions{Size: l.rect.Max}
}

func (l *Lamp) SetState(st bool) {
	l.state = st
}

func (l *Lamp) GetState() bool {
	return l.state
}

func (l *Lamp) Toggle() {
	l.state = !l.state
}

func (l *Lamp) getBodyColor() color.NRGBA {
	if l.state {
		return l.colorBody.Active
	}
	return l.colorBody.Normal
}

func (l *Lamp) getBorderColor() color.NRGBA {
	if l.state {
		return l.colorBorder.Active
	}
	return l.colorBorder.Normal
}
