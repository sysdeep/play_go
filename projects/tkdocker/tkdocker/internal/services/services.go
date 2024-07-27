package services

import "github.com/docker/docker/client"

type Services struct {
	Containers *ContainersService
}

func NewServices(docker_client *client.Client) *Services {

	return &Services{
		Containers: NewContainersService(docker_client),
	}
}
