package main

import (
	"gdocker/internal/ui"
	"gdocker/internal/ui/pages/containers"
	"gdocker/internal/ui/pages/page"
	"gdocker/internal/ui/pages/start"
	"log"
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
		window.Option(app.Title("GDocker"))
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
	running := true

	router := page.NewRouter()
	router.Register(0, start.NewStart())
	router.Register(1, containers.NewContainers())

	for {

		if !running {
			return nil
		}

		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:

			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			if closeButton.Clicked(gtx) {
				running = false
			}

			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return router.Layout(gtx, theme)
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return ui.SideNav{}.Layout(gtx, theme)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(theme, &closeButton, "Close").Layout(gtx)
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
