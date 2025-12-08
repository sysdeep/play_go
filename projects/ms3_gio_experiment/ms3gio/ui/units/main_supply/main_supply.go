package mainsupply

import (
	"image"
	"image/color"
	"ms3gio/internal/logic/models"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type MainSupply struct {
	pos   image.Point
	model *models.MainSupply
	max   image.Point
}

func New(model *models.MainSupply, pos image.Point) *MainSupply {
	return &MainSupply{
		pos:   pos,
		model: model,

		max: image.Pt(32, 64),
	}
}

func (s *MainSupply) Layout(gtx layout.Context) layout.Dimensions {

	defer op.Offset(s.pos).Push(gtx.Ops).Pop()

	body := clip.Rect{
		Min: image.Pt(0, 0),
		Max: s.max,
	}.Op()

	paint.FillShape(gtx.Ops, s.getColor(), body)

	return layout.Dimensions{Size: s.max}
}

func (d *MainSupply) getColor() color.NRGBA {
	return color.NRGBA{R: 100, G: 100, B: 100, A: 255}

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
