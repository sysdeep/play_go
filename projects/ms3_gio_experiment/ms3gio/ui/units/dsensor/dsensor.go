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
	max   image.Point
}

func New(model *models.DSensor, size int, pos image.Point) *DSensor {
	return &DSensor{
		pos:   pos,
		model: model,
		max:   image.Pt(size, size),
	}
}

func (s *DSensor) Layout(gtx layout.Context) layout.Dimensions {

	defer op.Offset(s.pos).Push(gtx.Ops).Pop()

	circle := clip.Ellipse{
		Min: image.Pt(0, 0),
		Max: s.max,
	}.Op(gtx.Ops)

	paint.FillShape(gtx.Ops, s.getColor(), circle)

	return layout.Dimensions{Size: s.max}
}

func (d *DSensor) getColor() color.NRGBA {
	if d.model.IsBlock {
		return color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	}

	if d.model.IsError {
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}

	if d.model.IsState {
		return color.NRGBA{R: 0, G: 255, B: 40, A: 255}
	}

	return color.NRGBA{R: 0, G: 100, B: 40, A: 255}
}
