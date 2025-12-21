package bs

import (
	"giobs/bs/label"
	"giobs/bs/theme"

	"gioui.org/font"
)

func H1(th *theme.Theme, txt string) label.LabelStyle {
	label := label.Label(th, th.TextSize*96.0/16.0, txt)
	label.Font.Weight = font.Light
	return label
}

func H2(th *theme.Theme, txt string) label.LabelStyle {
	label := label.Label(th, th.TextSize*60.0/16.0, txt)
	label.Font.Weight = font.Light
	return label
}

func H3(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*48.0/16.0, txt)
}

func H4(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*34.0/16.0, txt)
}

func H5(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*24.0/16.0, txt)
}

func H6(th *theme.Theme, txt string) label.LabelStyle {
	label := label.Label(th, th.TextSize*20.0/16.0, txt)
	label.Font.Weight = font.Medium
	return label
}

func Subtitle1(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*16.0/16.0, txt)
}

func Subtitle2(th *theme.Theme, txt string) label.LabelStyle {
	label := label.Label(th, th.TextSize*14.0/16.0, txt)
	label.Font.Weight = font.Medium
	return label
}

func Body1(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize, txt)
}

func Body2(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*14.0/16.0, txt)
}

func Caption(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*12.0/16.0, txt)
}

func Overline(th *theme.Theme, txt string) label.LabelStyle {
	return label.Label(th, th.TextSize*10.0/16.0, txt)
}
