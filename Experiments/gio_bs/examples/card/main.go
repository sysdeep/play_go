package main

import (
	"giobs/bs"
	"giobs/bs/card"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

type C = layout.Context
type D = layout.Dimensions

func main() {

	go func() {
		window := new(app.Window)
		window.Option(app.Title("Card"))
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
	bsTheme := bs.NewTheme()

	var ops op.Ops

	// var list = layout.List{}

	// var tabsList = []*tabs.Tab{}
	// for i := 1; i <= 10; i++ {
	// 	titleText := fmt.Sprintf("Tab %d", i)
	// 	tabContent := bs.H3(bsTheme, titleText)
	// 	tabsList = append(tabsList,
	// 		tabs.NewTab(bsTheme, titleText, func(gtx layout.Context) layout.Dimensions {
	// 			return tabContent.Layout(gtx)
	// 		}),
	// 	)
	// }

	// var tabs = tabs.NewTabs(bsTheme, tabsList...)

	var lbl = bs.Body1(bsTheme, "Hello!!!!\nlalala\nssss")
	var lblFooter = bs.Body1(bsTheme, "Footer")

	var card = card.NewCard(bsTheme, lbl.Layout, card.WithCardFooter(lblFooter.Layout))

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
				Spacing: layout.SpaceAround,
			}.Layout(gtx,

				layout.Rigid(func(gtx C) D {

					return layout.Flex{
						Axis:    layout.Horizontal,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return card.Layout(gtx)
						}),
					)

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
