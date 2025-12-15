package main

import (
	"image"
	"log"
	"ms3gio/internal/logic/models"
	mainsupply "ms3gio/ui/units/main_supply"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Title("DSensor"))
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

	var stateButton widget.Clickable
	var blockButton widget.Clickable
	var errorButton widget.Clickable

	sensorModel := &models.MainSupply{}
	sensorView1 := mainsupply.New(sensorModel, image.Pt(0, 0))
	sensorView2 := mainsupply.New(sensorModel, image.Pt(0, 0))
	sensorView3 := mainsupply.New(sensorModel, image.Pt(0, 0))

	// popupVisible := true

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// if stateButton.Clicked(gtx) {
			// 	popupVisible = true
			// }

			// if blockButton.Clicked(gtx) {
			// 	sensorModel.IsBlock = !sensorModel.IsBlock
			// }

			// if errorButton.Clicked(gtx) {
			// 	sensorModel.IsError = !sensorModel.IsError
			// }

			// draw
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				// scene
				// layout.Rigid(func(gtx C) D {
				// 	return Popup(&popupVisible).Layout(gtx, sensorView1.Layout)
				// }),
				layout.Rigid(sensorView1.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    unit.Dp(25),
						Bottom: unit.Dp(25),
						Right:  unit.Dp(35),
						Left:   unit.Dp(35),
					}.Layout(gtx, sensorView2.Layout)
				}),

				layout.Flexed(1.0,
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Horizontal,
						Alignment: layout.End,
						Spacing:   layout.SpaceAround,
					}.Layout(gtx, layout.Rigid(sensorView3.Layout))
				}),
				layout.Flexed(1.0,
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
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
							btn := material.Button(theme, &stateButton, "State")
							return btn.Layout(gtx)
						})

					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {

						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}

						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(theme, &blockButton, "Block")
							return btn.Layout(gtx)
						})

					},
				),

				layout.Rigid(
					func(gtx layout.Context) layout.Dimensions {

						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}

						return margins.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							btn := material.Button(theme, &errorButton, "Error")
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
