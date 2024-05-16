package webserver

import (
	"hdu/internal/logger"
	"hdu/internal/webserver/handlers"

	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Webserver struct {
	e *echo.Echo
	// docker *client.Client
}

func NewWebserver(docker *client.Client, logger *logger.Logger) *Webserver {

	e := echo.New()

	e.Static("/static", "public")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	// setup custom renderer
	tplr := NewTemplater()

	// prev templates
	// template_files := makeTemplatesList("views")
	// t := &Template{
	// 	// templates: template.Must(template.ParseGlob("views/*.html")),
	// 	templates: template.Must(template.ParseFiles(template_files...)),
	// }

	// e.Renderer = t
	e.Renderer = tplr

	// setup custom error renderer
	e.HTTPErrorHandler = customHTTPErrorHandler

	hndls := handlers.NewHandlers(docker, logger)

	e.GET("/", hndls.MainPage)
	e.GET("/containers/:id", hndls.ContainerPage)
	e.GET("/containers", hndls.ContainersPage)

	// volumes
	e.GET("/volumes/:name", hndls.VolumePage)
	e.GET("/volumes/actions/prune", hndls.ActionVolumesPrune)
	e.GET("/volumes/actions/remove/:name", hndls.ActionVolumeRemove)
	e.GET("/volumes", hndls.VolumesPage)

	// images
	e.GET("/images/:id", hndls.ImagePage)
	e.GET("/images", hndls.ImagesPage)
	// e.GET("/qqq", func(c echo.Context) error {

	// 	// return c.Render(200, "aaa", 0)
	// 	return c.Render(200, "aaa.html", 0)
	// })

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() {
	w.e.Logger.Fatal(w.e.Start("localhost:1313"))
}
