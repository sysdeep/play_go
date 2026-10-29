package button

import (
	"giobs/bs/theme"
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

	Th *theme.Theme
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

// OptWidget interface
func (b *Button) SetClass(className theme.ClassName) {
	switch className {

	case theme.Primary:
		b.Background = b.Th.Palette.Primary.Fg

	case theme.Secondary:
		b.Background = b.Th.Palette.Secondary.Fg

	case theme.Warning:
		b.Background = b.Th.Palette.Warning.Fg

	case theme.Success:
		b.Background = b.Th.Palette.Success.Fg

	case theme.Danger:
		b.Background = b.Th.Palette.Danger.Fg

	case theme.Light:
		b.Background = b.Th.Palette.Light.Fg

	case theme.Dark:
		b.Background = b.Th.Palette.Dark.Fg

	default:
		b.Background = b.Th.Palette.Primary.Fg
	}
}

// var stylesMap = map[theme.ClassName]string{
// 	theme.Primary: th.Palette.Primary.Fg
// }
