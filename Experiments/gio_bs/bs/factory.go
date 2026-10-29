package bs

import (
	"giobs/bs/button"
	"giobs/bs/theme"

	"gioui.org/layout"
	"gioui.org/widget"
)

type BSFactory struct {
	th *theme.Theme
}

func NewBSFactory(th *theme.Theme) *BSFactory {
	return &BSFactory{
		th: th,
	}
}

func (f *BSFactory) NewButton(txt string) button.Button {
	model := widget.Clickable{}
	b := button.Button{
		Text:         txt,
		Color:        f.th.Palette.ContrastFg,
		CornerRadius: 8,
		Background:   f.th.Palette.ContrastBg,
		TextSize:     f.th.TextSize * 14.0 / 16.0,
		Inset: layout.Inset{
			Top: 10, Bottom: 10,
			Left: 12, Right: 12,
		},
		Model:  &model,
		Shaper: f.th.Shaper,
	}
	b.Font.Typeface = f.th.Face
	// TODO: for debug
	b.Background = f.th.Palette.Primary.Fg
	return b
}
