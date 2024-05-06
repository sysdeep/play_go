package handlers

import (
	"hdu/internal/logger"

	"github.com/docker/docker/client"
)

type Handlers struct {
	docker_client *client.Client
	logger        *logger.Logger
}

func NewHandlers(docker_client *client.Client, logger *logger.Logger) *Handlers {

	return &Handlers{
		docker_client: docker_client,
		logger:        logger,
	}
}
