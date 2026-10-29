package bs

import (
	"giobs/bs/button"
	"giobs/bs/label"
	"giobs/bs/theme"

	"gioui.org/font"
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

func (f *BSFactory) NewButton(txt string, opts ...Opt) button.Button {
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
		Th:     f.th,
	}
	b.Font.Typeface = f.th.Face
	// TODO: for debug
	b.Background = f.th.Palette.Primary.Fg

	// b.SetClass(theme.Warning)
	for _, opt := range opts {
		opt(&b)
	}

	return b
}

func (f *BSFactory) NewLabel(txt string, opts ...Opt) label.LabelStyle {
	b := label.Label(f.th, f.th.TextSize, txt)

	// b.SetClass(theme.Warning)
	for _, opt := range opts {
		opt(&b)
	}

	return b
}

func (f *BSFactory) NewH6Label(txt string, opts ...Opt) label.LabelStyle {
	label := label.Label(f.th, f.th.TextSize*20.0/16.0, txt)
	label.Font.Weight = font.Medium

	for _, opt := range opts {
		opt(&label)
	}

	return label
}
