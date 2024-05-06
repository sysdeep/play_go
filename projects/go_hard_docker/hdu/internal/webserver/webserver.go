package webserver

import (
	"hdu/internal/logger"
	"hdu/internal/webserver/handlers"
	"html/template"

	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

type Webserver struct {
	e *echo.Echo
	// docker *client.Client
}

func NewWebserver(docker *client.Client, logger *logger.Logger) *Webserver {

	template_files := makeTemplatesList("views")

	t := &Template{
		// templates: template.Must(template.ParseGlob("views/*.html")),
		templates: template.Must(template.ParseFiles(template_files...)),
	}

	e := echo.New()
	e.Renderer = t

	hndls := handlers.NewHandlers(docker, logger)

	e.GET("/", hndls.MainPage)
	e.GET("/containers/:id", hndls.ContainerPage)
	e.GET("/containers", hndls.ContainersPage)
	e.GET("/volumes", hndls.VolumesPage)

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() {
	w.e.Logger.Fatal(w.e.Start("localhost:1313"))
}
