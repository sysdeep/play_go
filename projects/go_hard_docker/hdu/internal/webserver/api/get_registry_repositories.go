package api

import (
	"hdu/internal/registry_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

// handler
func (h *Api) GetRegistryRepositories(c echo.Context) error {

	catalog, err := h.registry_client.GetCatalog(10)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newRegistryRepositoryResponse(&catalog))
}

type registryRepositoriesResponse struct {
	Repositories []string `json:"repositories"`
}

func newRegistryRepositoryResponse(model *registry_client.Catalog) registryRepositoriesResponse {
	return registryRepositoriesResponse{
		Repositories: model.Repositories,
	}
}
