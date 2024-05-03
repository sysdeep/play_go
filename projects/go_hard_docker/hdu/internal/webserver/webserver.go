package webserver

import (
	"context"
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

	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()
	e.Renderer = t

	// test EP
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.GET("/hello", func(c echo.Context) error {
		// Получение списка запуцщенных контейнеров(docker ps)
		containers, err := docker.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			panic(err)
		}
		return c.Render(http.StatusOK, "hello", containers)
	})

	e.GET("/containers", func(c echo.Context) error {
		// Получение списка запуцщенных контейнеров(docker ps)
		containers, err := docker.ContainerList(context.Background(), container.ListOptions{All: true})
		if err != nil {
			panic(err)
		}
		return c.Render(http.StatusOK, "containers", containers)
	})

	return &Webserver{
		e: e,
	}
}

func (w *Webserver) Start() {
	w.e.Logger.Fatal(w.e.Start("localhost:1313"))
}
