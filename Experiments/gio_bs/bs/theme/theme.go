package theme

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
)

// Palette contains the minimal set of colors that a widget may need to
// draw itself.
type Palette struct {
	// Bg is the background color atop which content is currently being
	// drawn.
	Bg color.NRGBA

	// Fg is a color suitable for drawing on top of Bg.
	Fg color.NRGBA

	// ContrastBg is a color used to draw attention to active,
	// important, interactive widgets such as buttons.
	ContrastBg color.NRGBA

	// ContrastFg is a color suitable for content drawn on top of
	// ContrastBg.
	ContrastFg color.NRGBA

	// new bs
	BodyColor color.NRGBA
	BodyBg    color.NRGBA

	SecondaryColor color.NRGBA
	SecondaryBg    color.NRGBA

	TertiaryColor color.NRGBA
	TertiaryBg    color.NRGBA

	Emphasis color.NRGBA
	Border   color.NRGBA

	Primary   ActionColor
	Secondary ActionColor
	Success   ActionColor
	Danger    ActionColor
	Info      ActionColor
	Warning   ActionColor
	Light     ActionColor
	Dark      ActionColor
}

// Theme holds the general theme of an app or window. Different top-level
// windows should have different instances of Theme (with different Shapers;
// see the godoc for [text.Shaper]), though their other fields can be equal.
type Theme struct {
	Shaper *text.Shaper
	Palette
	TextSize unit.Sp
	Icon     struct {
		CheckBoxChecked   *widget.Icon
		CheckBoxUnchecked *widget.Icon
		RadioChecked      *widget.Icon
		RadioUnchecked    *widget.Icon
	}
	// Face selects the default typeface for text.
	Face font.Typeface

	// FingerSize is the minimum touch target size.
	FingerSize unit.Dp
}

// NewTheme constructs a theme (and underlying text shaper).
func NewTheme() *Theme {
	t := &Theme{Shaper: &text.Shaper{}}
	t.Palette = Palette{
		Fg:         rgb(0x000000),
		Bg:         rgb(0xffffff),
		ContrastBg: rgb(0x3f51b5),
		ContrastFg: rgb(0xffffff),

		// new
		BodyColor: rgb(0x212529),
		BodyBg:    rgb(0xfff),

		// SecondaryColor: rgb(0x0d6efd),
		SecondaryColor: color.NRGBA{A: 200, R: 33, B: 37, G: 41}, //rgba(33, 37, 41, 0.75);
		SecondaryBg:    rgb(0xe9ecef),

		// TertiaryColor: rgb(0x0d6efd),
		TertiaryColor: color.NRGBA{R: 33, G: 37, B: 41, A: 100}, //rgba(33, 37, 41, 0.5);
		TertiaryBg:    rgb(0xf8f9fa),

		Emphasis: rgb(0x000),
		Border:   rgb(0xdee2e6),

		Primary: ActionColor{
			Fg:     rgb(0x0d6efd),
			Bg:     rgb(0xcfe2ff),
			Border: rgb(0x9ec5fe),
			Text:   rgb(0x052c65),
		},

		Secondary: ActionColor{
			Fg:     rgb(0x6c757d),
			Bg:     rgb(0xe2e3e5),
			Border: rgb(0xc4c8cb),
			Text:   rgb(0x2b2f32),
		},

		Success: ActionColor{
			Fg: rgb(0x198754),
		},
		Danger: ActionColor{
			Fg: rgb(0xdc3545),
		},
		Info: ActionColor{
			Fg: rgb(0x0dcaf0),
		},
		Warning: ActionColor{
			Fg: rgb(0xffc107),
		},
		Light: ActionColor{
			Fg: rgb(0xf8f9fa),
		},
		Dark: ActionColor{
			Fg: rgb(0x212529),
		},
	}
	t.TextSize = 16

	// t.Icon.CheckBoxChecked = mustIcon(widget.NewIcon(icons.ToggleCheckBox))
	// t.Icon.CheckBoxUnchecked = mustIcon(widget.NewIcon(icons.ToggleCheckBoxOutlineBlank))
	// t.Icon.RadioChecked = mustIcon(widget.NewIcon(icons.ToggleRadioButtonChecked))
	// t.Icon.RadioUnchecked = mustIcon(widget.NewIcon(icons.ToggleRadioButtonUnchecked))

	// 38dp is on the lower end of possible finger size.
	t.FingerSize = 38

	return t
}

func (t Theme) WithPalette(p Palette) Theme {
	t.Palette = p
	return t
}

func mustIcon(ic *widget.Icon, err error) *widget.Icon {
	if err != nil {
		panic(err)
	}
	return ic
}

func rgb(c uint32) color.NRGBA {
	return argb(0xff000000 | c)
}

func argb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}
