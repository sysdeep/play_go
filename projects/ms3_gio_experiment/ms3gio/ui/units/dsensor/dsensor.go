package dsensor

import (
	"image"
	"image/color"
	"ms3gio/internal/logic/models"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type DSensor struct {
	pos   image.Point
	model *models.DSensor
}

func New(model *models.DSensor, pos image.Point) *DSensor {
	return &DSensor{
		pos:   pos,
		model: model,
	}
}

func (s *DSensor) Layout(gtx layout.Context) layout.Dimensions {

	defer op.Offset(s.pos).Push(gtx.Ops).Pop()

	circle := clip.Ellipse{
		// Hard coding the x coordinate. Try resizing the window
		// Min: image.Pt(80, 0),
		// Max: image.Pt(320, 240),
		// Soft coding the x coordinate. Try resizing the window
		Min: image.Pt(0, 0),
		Max: image.Pt(32, 32),
	}.Op(gtx.Ops)

	o_color := s.getColor()
	// if boiling {
	// 	o_color = color.NRGBA{G: 200, A: 255}
	// }
	paint.FillShape(gtx.Ops, o_color, circle)

	d := image.Point{Y: 32, X: 32}

	return layout.Dimensions{Size: d}
}

func (d *DSensor) getColor() color.NRGBA {
	if d.model.IsBlock {
		return color.NRGBA{R: 50, G: 50, B: 50, A: 255}
	}

	if d.model.IsError {
		return color.NRGBA{R: 250, G: 0, B: 0, A: 255}
	}

	if d.model.IsState {
		return color.NRGBA{R: 0, G: 250, B: 0, A: 255}
	}

	return color.NRGBA{R: 0, G: 50, B: 0, A: 255}
}
