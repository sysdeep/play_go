package webserver

import (
	"context"
	"hdu/internal/webserver/handlers"
	"html/template"
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
)

type Webserver struct {
	e *echo.Echo
	// docker *client.Client
}

func NewWebserver(docker *client.Client) *Webserver {

	template_files := makeTemplatesList("views")

	t := &Template{
		// templates: template.Must(template.ParseGlob("views/*.html")),
		templates: template.Must(template.ParseFiles(template_files...)),
	}

	e := echo.New()
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		// Получение списка запуцщенных контейнеров(docker ps)
		containers, err := docker.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			panic(err)
		}
		return c.Render(http.StatusOK, "main", containers)
	})

	hndls := handlers.NewHandlers(docker)

	e.GET("/containers", hndls.ContainersPage)

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() {
	w.e.Logger.Fatal(w.e.Start("localhost:1313"))
}
