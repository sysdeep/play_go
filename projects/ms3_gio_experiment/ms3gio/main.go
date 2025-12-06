package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("Ms3 gio example"))
		window.Option(app.Size(unit.Dp(800), unit.Dp(600)))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()

	var ops op.Ops

	var closeButton widget.Clickable
	var boiling bool

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if closeButton.Clicked(gtx) {
				boiling = !boiling
			}

			// draw
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				// circle
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						circle := clip.Ellipse{
							// Hard coding the x coordinate. Try resizing the window
							// Min: image.Pt(80, 0),
							// Max: image.Pt(320, 240),
							// Soft coding the x coordinate. Try resizing the window
							Min: image.Pt(gtx.Constraints.Max.X/2-120, 0),
							Max: image.Pt(gtx.Constraints.Max.X/2+120, 240),
						}.Op(gtx.Ops)

						o_color := color.NRGBA{R: 200, A: 255}
						if boiling {
							o_color = color.NRGBA{G: 200, A: 255}
						}
						paint.FillShape(gtx.Ops, o_color, circle)

						d := image.Point{Y: 400}

						return layout.Dimensions{Size: d}
					},
				),

				// text
				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {
						// Define an large label with an appropriate text:
						title := material.H1(theme, fmt.Sprintf("Hello, Gio: %t", boiling))

						// Change the color of the label
						maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
						title.Color = maroon

						// Change the position of the label.
						title.Alignment = text.Middle

						// Draw the label to the graphics context.
						return title.Layout(gtx)
					},
				),

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
							btn := material.Button(theme, &closeButton, "Toggle")
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
