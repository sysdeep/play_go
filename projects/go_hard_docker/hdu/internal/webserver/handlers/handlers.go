package handlers

import "github.com/docker/docker/client"

type Handlers struct {
	docker_client *client.Client
}

func NewHandlers(docker_client *client.Client) *Handlers {

	return &Handlers{docker_client}
}
