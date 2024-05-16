package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// models
type secretListModel struct {
}
type secretsPageModel struct {
	Secrets []secretListModel
	Total   int
}

// handler
func (h *Handlers) SecretsPage(c echo.Context) error {
	// volumes_data, err := h.docker_client.VolumeList(context.Background(), volume.ListOptions{})
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusInternalServerError, err)
	// }
	//
	// var volumes []volumeListModel
	// for _, v := range volumes_data.Volumes {
	// 	volumes = append(volumes, make_volume_list_model(v))
	// }
	//
	// sort.SliceStable(volumes, func(i, j int) bool {
	// 	return volumes[i].Name < volumes[j].Name
	// })

	response := secretsPageModel{
		Secrets: *new([]secretListModel),
		Total:   0,
	}

	return c.Render(http.StatusOK, "secrets.html", response)
}
