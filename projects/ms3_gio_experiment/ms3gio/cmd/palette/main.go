package main

import (
	"image"
	"image/color"
	"log"
	"ms3gio/ui/palette"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
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

	// r := utils.ColorFrom(utils.BlueGray, utils.P50)
	// fmt.Println(r)

	// r = utils.ColorFrom(utils.BlueGray, utils.P100)
	// fmt.Println(r)
}

func run(window *app.Window) error {
	var list = layout.List{Axis: layout.Vertical}

	var ops op.Ops

	// var stateButton widget.Clickable
	// var blockButton widget.Clickable
	// var errorButton widget.Clickable

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

				// 				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// return list.Layout(gtx, len(palette.Colors), func(gtx layout.Context, i int) layout.Dimensions {

				// })
				// 				},

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					cc := palette.Colors[0:]
					return list.Layout(gtx, len(cc), func(gtx layout.Context, i int) layout.Dimensions {
						return makeRow(gtx, cc[i])
					})
				}),

				// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// 	cc := palette.Colors[6:8]
				// 	return list.Layout(gtx, len(cc), func(gtx layout.Context, i int) layout.Dimensions {
				// 		return makeRow(gtx, cc[i])
				// 	})
				// }),

				// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// 	return makeRow(gtx, palette.BlueGray)
				// }),

				// layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				// 	return list.Layout(gtx, len(palette.Scales), func(gtx layout.Context, i int) layout.Dimensions {
				// 		col := palette.Color(palette.CoolGray, palette.Scales[i])
				// 		return ColorBox(gtx, image.Pt(20, 100), col)
				// 	})
				// }),

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

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func makeRow(gtx layout.Context, paletteColor int) layout.Dimensions {
	var list = layout.List{}

	return list.Layout(gtx, len(palette.Scales), func(gtx layout.Context, i int) layout.Dimensions {
		col := palette.Color(paletteColor, palette.Scales[i])
		return ColorBox(gtx, image.Pt(32, 32), col)
	})
}
