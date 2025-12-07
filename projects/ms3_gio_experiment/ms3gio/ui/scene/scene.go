package scene

import (
	"image"
	"image/color"
	"ms3gio/internal/logic/project"
	"ms3gio/ui/units/section"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Scene struct {
	project *project.Project

	// units
	section *section.Section
}

func New(project *project.Project) *Scene {
	return &Scene{
		project: project,
		section: section.New(project.Section),
	}
}

func (s *Scene) Layout(gtx layout.Context) layout.Dimensions {
	rect := clip.Rect{
		Max: gtx.Constraints.Max,
	}.Op()

	// circle := clip.Ellipse{
	// 	// Hard coding the x coordinate. Try resizing the window
	// 	// Min: image.Pt(80, 0),
	// 	// Max: image.Pt(320, 240),
	// 	// Soft coding the x coordinate. Try resizing the window
	// 	Min: image.Pt(gtx.Constraints.Max.X/2-120, 0),
	// 	Max: image.Pt(gtx.Constraints.Max.X/2+120, 240),
	// }.Op(gtx.Ops)

	o_color := color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	// if boiling {
	// 	o_color = color.NRGBA{G: 200, A: 255}
	// }
	// paint.FillShape(gtx.Ops, o_color, circle)
	paint.FillShape(gtx.Ops, o_color, rect)

	d := image.Point{Y: 400}

	s.section.Layout(gtx)

	return layout.Dimensions{Size: d}
}
