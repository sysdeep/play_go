package services

import (
	"context"

	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type VolumesService struct {
	docker_client *client.Client
}

func NewVolumesService(docker_client *client.Client) *VolumesService {
	return &VolumesService{docker_client: docker_client}
}

func (cs *VolumesService) GetAll() ([]VolumeListModel, error) {
	var result []VolumeListModel
	volumes_list, err := cs.docker_client.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return result, err
	}

	for _, c := range volumes_list.Volumes {
		result = append(result, make_volume_list_model(c))
	}

	return result, nil
}

// models ---------------------------------------------------------------------

type VolumeListModel struct {
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

	StackName string
}

const volume_stack_label = "com.docker.stack.namespace"

func make_volume_list_model(data *volume.Volume) VolumeListModel {

	stack_name := ""
	if stack_name_labeled, ok := data.Labels[volume_stack_label]; ok {
		stack_name = stack_name_labeled
	}

	return VolumeListModel{
		Name:       data.Name,
		CreatedAt:  data.CreatedAt,
		Driver:     data.Driver,
		Mountpoint: data.Mountpoint,
		StackName:  stack_name,
	}

}
