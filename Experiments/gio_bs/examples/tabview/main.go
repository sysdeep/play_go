package main

import (
	"giobs/bs"
	"giobs/bs/button"
	"giobs/bs/tabview"
	"giobs/bs/theme"
	"image"
	"image/color"
	"log"
	"math/rand/v2"
	"os"

	"gioui.org/app"
	"gioui.org/font"
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
	// themeMateria := material.NewTheme()
	bsTheme := bs.NewTheme()

	var ops op.Ops

	var stateButton widget.Clickable
	var blockButton widget.Clickable
	var errorButton widget.Clickable

	var mainColor = color.NRGBA{A: 255}

	var flexLayout = layout.Flex{
		// Vertical alignment, from top to bottom
		Axis: layout.Vertical,
		// Empty space is left at the start, i.e. at the top
		Spacing: layout.SpaceStart,
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
			flexLayout.Layout(gtx,

				layout.Rigid(func(gtx C) D {

					return tabview.NewTabView(layout.Horizontal, buildTabItems()...).Layout(gtx, bsTheme)
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

func ColorBox(gtx layout.Context, size image.Point, color color.NRGBA) layout.Dimensions {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return layout.Dimensions{Size: size}
}

func buildTabItems() []*tabview.TabItem {
	inset := layout.Inset{
		Left:   unit.Dp(12),
		Right:  unit.Dp(12),
		Top:    unit.Dp(8),
		Bottom: unit.Dp(8),
	}

	var tabItems []*tabview.TabItem
	tabItems = append(tabItems, makeTabItem(inset, "Tab 1", func(gtx C, th *theme.Theme) D {
		return layoutTab(gtx, th, "Tab one")
	}))

	tabItems = append(tabItems, makeTabItem(inset, "A long tab name", func(gtx C, th *theme.Theme) D {
		return layoutTab(gtx, th, "Tab two")
	}))

	tabItems = append(tabItems, makeTabItem(inset, "Tab 3", func(gtx C, th *theme.Theme) D {
		return layoutTab(gtx, th, "Tab three")
	}))

	tabItems = append(tabItems, makeTabItem(inset, "Tab 4", func(gtx C, th *theme.Theme) D {
		return layoutTab(gtx, th, "Tab four")
	}))

	tabItems = append(tabItems, makeTabItem(inset, "Tab 5", func(gtx C, th *theme.Theme) D {
		return layoutTab(gtx, th, "Tab five")
	}))

	return tabItems

}

func makeTabItem(inset layout.Inset, title string, wgt func(gtx C, th *theme.Theme) D) *tabview.TabItem {
	thm := material.NewTheme()
	return &tabview.TabItem{
		Title: func(gtx C, th *theme.Theme) D {
			label := material.Label(thm, th.TextSize, title)
			label.Font.Weight = font.Medium
			return label.Layout(gtx)
		},
		Widget: wgt,
		Inset:  inset,
	}
}

func layoutTab(gtx C, th *theme.Theme, content string) D {
	thm := material.NewTheme()
	return layout.Center.Layout(gtx, func(gtx C) D {
		label := material.Label(thm, th.TextSize*0.9, content)
		label.Font.Typeface = font.Typeface("Go Mono")
		return label.Layout(gtx)
	})
}
