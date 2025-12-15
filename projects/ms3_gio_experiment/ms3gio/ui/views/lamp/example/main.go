package main

import (
	"image"
	"image/color"
	"log"
	"ms3gio/ui/palette"
	"ms3gio/ui/views/lamp"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Lamp"))
		window.Option(app.Size(unit.Dp(800), unit.Dp(600)))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func makeColorsPair(color int) (lamp.LampColor, lamp.LampColor) {
	clBody := lamp.LampColor{
		Active: palette.Color(color, palette.P400),
		Normal: palette.Color(color, palette.P600),
	}

	clBorder := lamp.LampColor{
		Active: palette.Color(color, palette.P600),
		Normal: palette.Color(color, palette.P400),
	}

	return clBody, clBorder
}

func run(window *app.Window) error {
	theme := material.NewTheme()

	var ops op.Ops

	var stateButton widget.Clickable

	colors := palette.Colors[:]

	lamps := []*lamp.Lamp{}
	for _, color := range colors {
		body, border := makeColorsPair(color)
		lamps = append(lamps, lamp.New(32, body, border))
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if stateButton.Clicked(gtx) {
				for _, lamp := range lamps {
					lamp.Toggle()
				}
			}

			// draw
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				// scene
				// layout.Rigid(sensorView.Layout),
				layout.Rigid(func(gtx C) D { return makeSceneLayout(gtx, lamps) }),

				// actions row
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {

						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}

						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(theme, &stateButton, "State")
							return btn.Layout(gtx)
						})

					},
				),

				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}

func makeSceneLayout(gtx C, items []*lamp.Lamp) D {
	var list = layout.List{}
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			return ColorBox(gtx, image.Pt(1200, 200), palette.Color(palette.Gray, palette.P300))
		}),
		layout.Stacked(func(gtx C) D {
			return list.Layout(gtx, len(items), func(gtx C, i int) D {
				return items[i].Layout(gtx)
			})
		}),
	)
}

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
