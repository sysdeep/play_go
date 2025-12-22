package card

import (
	"giobs/bs/theme"
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

type Card struct {
	th      *theme.Theme
	content layout.Widget
	footer  layout.Widget
}

type CardOption = func(c *Card)

func NewCard(th *theme.Theme, content layout.Widget, options ...CardOption) *Card {
	card := &Card{
		th:      th,
		content: content,
	}

	for _, o := range options {

		o(card)
	}

	return card
}

func WithCardFooter(content layout.Widget) CardOption {
	return func(c *Card) {
		c.footer = content
	}
}

func (c *Card) Layout(gtx layout.Context) layout.Dimensions {

	var contentDim layout.Dimensions
	var footerDim layout.Dimensions

	return layout.Stack{}.Layout(gtx,

		// content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(
				gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					contentDim = layout.UniformInset(unit.Dp(12)).Layout(gtx, c.content)
					return contentDim
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if c.footer != nil {
						footerDim = layout.UniformInset(unit.Dp(12)).Layout(gtx, c.footer)
						return footerDim
					}
					return layout.Dimensions{}
				}),
			)
		}),

		// border
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {

			maxDim := layout.Dimensions{
				Size: image.Point{
					X: contentDim.Size.X,
					Y: contentDim.Size.Y + footerDim.Size.Y,
				},
			}

			tabRect := image.Rect(0, 0, maxDim.Size.X, maxDim.Size.Y)
			const r = 6
			rrect := clip.UniformRRect(tabRect, r)
			paint.FillShape(
				gtx.Ops,
				c.th.BodyColor,
				clip.Stroke{
					Path:  rrect.Path(gtx.Ops),
					Width: 1,
				}.Op(),
			)

			return layout.Dimensions{
				Size: tabRect.Max,
			}
		}),
	)
}
