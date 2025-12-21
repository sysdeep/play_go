package main

import (
	"giobs/bs"
	"giobs/bs/button"
	"giobs/bs/label"
	"giobs/bs/theme"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
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

func run(window *app.Window) error {
	bsTheme := bs.NewTheme()

	var ops op.Ops

	var list = layout.List{}

	var labels = []label.LabelStyle{
		bs.H1(bsTheme, "H1"),
		bs.H2(bsTheme, "H2"),
		bs.H3(bsTheme, "H3"),
		bs.H4(bsTheme, "H4"),
		bs.H5(bsTheme, "H5"),
		bs.H6(bsTheme, "H6"),
		bs.Subtitle1(bsTheme, "Subtitle1"),
		bs.Subtitle2(bsTheme, "Subtitle2"),
		bs.Body1(bsTheme, "Body1"),
		bs.Body2(bsTheme, "Body2"),
		bs.Caption(bsTheme, "Caption"),
		bs.Overline(bsTheme, "Overline"),
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// draw
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(func(gtx C) D {
					return list.Layout(gtx, len(labels), func(gtx C, i int) D {

						margins := layout.Inset{
							Top:    unit.Dp(12),
							Bottom: unit.Dp(12),
							Right:  unit.Dp(12),
							Left:   unit.Dp(12),
						}

						return margins.Layout(gtx, func(gtx C) D {
							return labels[i].Layout(gtx)
						})

					})
				}),

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
