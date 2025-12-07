package modals

import (
	"image/color"
	"ms3gio/ui/widgets/card"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type MessageType string

const (
	MessageTypeInfo MessageType = "info"
	MessageTypeWarn MessageType = "warn"
	MessageTypeErr  MessageType = "err"
)

type Message struct {
	Title string
	Body  string
	Type  MessageType

	OKBtn widget.Clickable

	Th *material.Theme
}

// func (m *Message) Layout(gtx layout.Context, th *chapartheme.Theme) layout.Dimensions {
func (m *Message) Layout(gtx layout.Context) layout.Dimensions {
	// textColor := th.Palette.Fg
	textColor := color.NRGBA{R: 0xD1, G: 0x1E, B: 0x35, A: 0xFF}
	switch m.Type {
	case MessageTypeErr:
		textColor = color.NRGBA{R: 0xD1, G: 0x1E, B: 0x35, A: 0xFF}
	case MessageTypeInfo:
		textColor = color.NRGBA{R: 0x1D, G: 0xBF, B: 0xEC, A: 0xFF}
	case MessageTypeWarn:
		textColor = color.NRGBA{R: 0xFD, G: 0xB5, B: 0x0E, A: 0xFF}
	}

	marginTop := layout.Inset{Top: unit.Dp(90)}

	return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X /= 3
			return marginTop.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return card.Card{
					Th: m.Th,

					Title: m.Title,
					Body: func(gtx layout.Context) layout.Dimensions {
						lb := material.Label(m.Th, unit.Sp(14), m.Body)
						lb.Color = textColor
						return lb.Layout(gtx)
					},

					Actions: []card.Action{
						{
							Clickable: &m.OKBtn,
							Label:     "Ok",
							Fg:        color.NRGBA{R: 100, G: 100, B: 100, A: 0xFF},
							Bg:        color.NRGBA{R: 50, G: 50, B: 50, A: 0xFF},
							Float:     card.FloatRight,
						},
					},
				}.Layout(gtx)
			})
		})
	})
}
