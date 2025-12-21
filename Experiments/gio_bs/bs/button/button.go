package button

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

type Button struct {
	Text string
	// Color is the text color.
	Color        color.NRGBA
	Font         font.Font
	TextSize     unit.Sp
	Background   color.NRGBA
	CornerRadius unit.Dp
	Inset        layout.Inset
	Model        *widget.Clickable
	Shaper       *text.Shaper
}

func (b Button) Layout(gtx layout.Context) layout.Dimensions {
	return ButtonLayoutStyle{
		Background:   b.Background,
		CornerRadius: b.CornerRadius,
		Model:        b.Model,
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return b.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			colMacro := op.Record(gtx.Ops)
			paint.ColorOp{Color: b.Color}.Add(gtx.Ops)
			return widget.Label{Alignment: text.Middle}.
				Layout(gtx, b.Shaper, b.Font, b.TextSize, b.Text, colMacro.Stop())
		})
	})
}
