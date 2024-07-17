package main

import "github.com/pwiecz/go-fltk"

func main() {
	is_clicked := false

	win := fltk.NewWindow(400, 300)
	win.SetLabel("Main Window")
	btn := fltk.NewButton(160, 200, 80, 30, "Click")
	btn.SetCallback(func() {
		is_clicked = !is_clicked
		if is_clicked {
			btn.SetLabel("Clicked")
		} else {
			btn.SetLabel("Click")
		}
	})
	win.End()
	win.Show()
	fltk.Run()
}
