package api

import (
	"hdu/internal/logger"
	"hdu/internal/services"

	"github.com/docker/docker/client"
)

type Api struct {
	docker_client *client.Client
	logger        *logger.Logger
	services      *services.Services
}

func NewApi(docker_client *client.Client, services *services.Services, logger *logger.Logger) *Api {

	return &Api{
		docker_client: docker_client,
		logger:        logger,
		services:      services,
	}
}
