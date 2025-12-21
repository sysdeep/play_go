package label

import (
	"giobs/bs/internal/f32color"
	"giobs/bs/theme"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
)

func Label(th *theme.Theme, size unit.Sp, txt string) LabelStyle {
	l := LabelStyle{
		Text:           txt,
		Color:          th.Palette.Fg,
		SelectionColor: f32color.MulAlpha(th.Palette.ContrastBg, 0x60),
		TextSize:       size,
		Shaper:         th.Shaper,
	}
	l.Font.Typeface = th.Face
	return l
}

func (l LabelStyle) Layout(gtx layout.Context) layout.Dimensions {
	textColorMacro := op.Record(gtx.Ops)
	paint.ColorOp{Color: l.Color}.Add(gtx.Ops)
	textColor := textColorMacro.Stop()
	selectColorMacro := op.Record(gtx.Ops)
	paint.ColorOp{Color: l.SelectionColor}.Add(gtx.Ops)
	selectColor := selectColorMacro.Stop()

	if l.State != nil {
		if l.State.Text() != l.Text {
			l.State.SetText(l.Text)
		}
		l.State.Alignment = l.Alignment
		l.State.MaxLines = l.MaxLines
		l.State.Truncator = l.Truncator
		l.State.WrapPolicy = l.WrapPolicy
		l.State.LineHeight = l.LineHeight
		l.State.LineHeightScale = l.LineHeightScale
		return l.State.Layout(gtx, l.Shaper, l.Font, l.TextSize, textColor, selectColor)
	}
	tl := widget.Label{
		Alignment:       l.Alignment,
		MaxLines:        l.MaxLines,
		Truncator:       l.Truncator,
		WrapPolicy:      l.WrapPolicy,
		LineHeight:      l.LineHeight,
		LineHeightScale: l.LineHeightScale,
	}
	return tl.Layout(gtx, l.Shaper, l.Font, l.TextSize, l.Text, textColor)
}
