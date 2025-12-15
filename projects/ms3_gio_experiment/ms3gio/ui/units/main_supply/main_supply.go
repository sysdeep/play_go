package mainsupply

import (
	"image"
	"ms3gio/internal/logic/models"

	"gioui.org/layout"
)

type MainSupply struct {
	pos   image.Point
	model *models.MainSupply
	// Bounding image.Rectangle

	// components
	body    *body
	powerSW *powerSwitcher
}

func New(model *models.MainSupply, pos image.Point) *MainSupply {

	return &MainSupply{
		pos:   pos,
		model: model,
		// Bounding: image.Rect(0, 0, 48, 64),

		body:    newBody(),
		powerSW: newPowerSwitcher(image.Pt(4, 16)),
	}
}

func (s *MainSupply) Layout(gtx layout.Context) layout.Dimensions {

	if s.powerSW.Clicked() {
		s.model.PowerOn()
	}

	return layout.Stack{}.Layout(gtx,
		// Force widget to the same size as the second.
		layout.Expanded(s.body.Layout),
		layout.Stacked(s.powerSW.Layout),
	)
}

func (m *MainSupply) BoundingRect() image.Rectangle {
	return m.body.BoundingRect()
}
