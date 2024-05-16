package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// models
type configListModel struct {
}
type configsPageModel struct {
	Configs []configListModel
	Total   int
}

// handler
func (h *Handlers) ConfigsPage(c echo.Context) error {
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

	response := configsPageModel{
		Configs: *new([]configListModel),
		Total:   0,
	}

	return c.Render(http.StatusOK, "configs.html", response)
}
