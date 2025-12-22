package bs

import (
	"giobs/bs/button"
	"giobs/bs/theme"

	"gioui.org/layout"
	"gioui.org/widget"
)

func Button(th *theme.Theme, model *widget.Clickable, txt string) button.Button {
	b := button.Button{
		Text:         txt,
		Color:        th.Palette.ContrastFg,
		CornerRadius: 8,
		Background:   th.Palette.ContrastBg,
		TextSize:     th.TextSize * 14.0 / 16.0,
		Inset: layout.Inset{
			Top: 10, Bottom: 10,
			Left: 12, Right: 12,
		},
		Model:  model,
		Shaper: th.Shaper,
	}
	b.Font.Typeface = th.Face
	return b
}

func ButtonPrimary(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Primary.Fg
	return b
}

func ButtonSecondary(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Secondary.Fg
	return b
}

func ButtonSuccess(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Success.Fg
	return b
}

func ButtonDanger(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Danger.Fg
	return b
}

func ButtonInfo(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Info.Fg
	return b
}

func ButtonWarning(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Warning.Fg
	return b
}

func ButtonLight(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Light.Fg
	return b
}

func ButtonDark(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button {
	b := Button(th, buttonCli, txt)
	b.Background = th.Palette.Dark.Fg
	return b
}

// Primary
// Secondary
// Success
// Danger
// Warning
// Info
// Light
// Dark
