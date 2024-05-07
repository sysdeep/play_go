package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/volume"
	"github.com/labstack/echo/v4"
)

// models
type volumeListModel struct {
	Name string

	// cluster volume
	// ClusterVolume *ClusterVolume `json:"ClusterVolume,omitempty"`

	// Date/Time the volume was created.
	CreatedAt string

	// Name of the volume driver used by the volume.
	Driver string

	// User-defined key/value metadata.
	// Required: true
	// Labels map[string]string `json:"Labels"`

	// Mount path of the volume on the host.
	Mountpoint string

	// The driver specific options used when creating the volume.
	//
	// Required: true
	// Options map[string]string `json:"Options"`

	// The level at which the volume exists. Either `global` for cluster-wide,
	// or `local` for machine level.
	//
	// Required: true
	// Scope string `json:"Scope"`

	// Low-level details about the volume, provided by the volume driver.
	// Details are returned as a map with key/value pairs:
	// `{"key":"value","key2":"value2"}`.
	//
	// The `Status` field is optional, and is omitted if the volume driver
	// does not support this feature.
	//
	// Status map[string]interface{} `json:"Status,omitempty"`

	// usage data
	// UsageData *UsageData `json:"UsageData,omitempty"`

}

type volumesPageModel struct {
	Volumes []volumeListModel
}

// handler
func (h *Handlers) VolumesPage(c echo.Context) error {
	volumes_data, err := h.docker_client.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		panic(err)
	}

	var volumes []volumeListModel
	for _, v := range volumes_data.Volumes {
		volumes = append(volumes, make_volume_list_model(v))
	}

	response := volumesPageModel{
		Volumes: volumes,
	}

	return c.Render(http.StatusOK, "volumes", response)
}

func make_volume_list_model(data *volume.Volume) volumeListModel {
	fmt.Printf("%+v\n", data)
	return volumeListModel{
		Name:       data.Name,
		CreatedAt:  data.CreatedAt,
		Driver:     data.Driver,
		Mountpoint: data.Mountpoint,
	}
}
