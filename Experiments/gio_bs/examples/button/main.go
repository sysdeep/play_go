package main

import (
	"giobs/bs"
	"giobs/bs/button"
	"giobs/bs/theme"
	"image"
	"image/color"
	"log"
	"math/rand/v2"
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

type Btn = func(th *theme.Theme, buttonCli *widget.Clickable, txt string) button.Button
type BtnDef struct {
	Btn  Btn
	Name string
}

func run(window *app.Window) error {
	themeMateria := material.NewTheme()
	bsTheme := bs.NewTheme()

	var ops op.Ops

	var stateButton widget.Clickable
	var blockButton widget.Clickable
	var errorButton widget.Clickable

	var mainColor = color.NRGBA{A: 255}

	var list = layout.List{}

	var buttons = []BtnDef{
		{Btn: bs.ButtonPrimary, Name: "Primary"},
		{Btn: bs.ButtonSecondary, Name: "Secondary"},
		{Btn: bs.ButtonSuccess, Name: "Success"},
		{Btn: bs.ButtonDanger, Name: "Danger"},
		{Btn: bs.ButtonInfo, Name: "Info"},
		{Btn: bs.ButtonWarning, Name: "Warning"},
		{Btn: bs.ButtonLight, Name: "Light"},
		{Btn: bs.ButtonDark, Name: "Dark"},
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if stateButton.Clicked(gtx) {
				mainColor.R = uint8(rand.IntN(255))
			}

			if blockButton.Clicked(gtx) {
				mainColor.G = uint8(rand.IntN(255))
			}

			if errorButton.Clicked(gtx) {
				mainColor.B = uint8(rand.IntN(255))
			}

			// draw
			layout.Flex{
				// Vertical alignment, from top to bottom
				Axis: layout.Vertical,
				// Empty space is left at the start, i.e. at the top
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(func(gtx C) D {
					return list.Layout(gtx, len(buttons), func(gtx C, i int) D {

						margins := layout.Inset{
							Top:    unit.Dp(12),
							Bottom: unit.Dp(12),
							Right:  unit.Dp(12),
							Left:   unit.Dp(12),
						}

						return margins.Layout(gtx, func(gtx C) D {
							return buttons[i].Btn(bsTheme, &stateButton, buttons[i].Name).Layout(gtx)
						})

					})
				}),

				layout.Rigid(func(gtx C) D {
					return ColorBox(gtx, image.Pt(300, 300), mainColor)
				}),

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
							btn := bs.Button(bsTheme, &stateButton, "State")
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
							btn := bs.ButtonDanger(bsTheme, &stateButton, "State")
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
							btn := material.Button(themeMateria, &blockButton, "Block")
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
							btn := material.Button(themeMateria, &errorButton, "Error")
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

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}
