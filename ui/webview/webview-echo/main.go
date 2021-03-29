package main

// https://github.com/webview/webview

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webview/webview"
)

func startServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<html>
			<head><title>Example</title></head>
			<body>
				<a href="/page1">page1</a>
			</body>
		</html>
		`)
	})

	e.GET("/page1", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<html>
			<head><title>Example</title></head>
			<body>
				<a href="/">main</a>
			</body>
		</html>
		`)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func main() {

	go startServer()

	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800, 600, webview.HintNone)
	// w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	w.Navigate("http://localhost:1323")
	w.Run()
}
