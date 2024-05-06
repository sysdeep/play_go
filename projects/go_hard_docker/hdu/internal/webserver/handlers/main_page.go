package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type mainPageModel struct{}

func (h *Handlers) MainPage(c echo.Context) error {
	response := mainPageModel{}
	return c.Render(http.StatusOK, "main", response)
}
