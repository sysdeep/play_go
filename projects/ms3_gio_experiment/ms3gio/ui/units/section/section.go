package section

import (
	"image"
	"image/color"
	"ms3gio/internal/logic/models"
	"ms3gio/ui/units/dsensor"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Section struct {
	model *models.Section

	// units
	sensor1 *dsensor.DSensor
}

func New(model *models.Section) *Section {

	return &Section{
		model:   model,
		sensor1: dsensor.New(model.Sensor1, 24, image.Pt(40, 400)),
	}
}

func (s *Section) Layout(gtx layout.Context) layout.Dimensions {

	rect := clip.Rect{
		Max: image.Point{500, 500},
	}.Op()

	o_color := color.NRGBA{R: 10, G: 100, B: 100, A: 255}
	paint.FillShape(gtx.Ops, o_color, rect)

	d := image.Point{Y: 400}

	s.sensor1.Layout(gtx)

	return layout.Dimensions{Size: d}
}
