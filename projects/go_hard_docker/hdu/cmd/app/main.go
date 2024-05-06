package main

import (
	"context"
	"fmt"
	"hdu/internal/logger"
	"hdu/internal/webserver"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	log := logger.NewLogger()
	log.Info("start")

	// test docker

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Получение списка запуцщенных контейнеров(docker ps)
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	// Вывод всех идентификаторов контейнеров
	for _, container := range containers {
		// fmt.Println(container.ID)
		fmt.Printf("%s %s (status: %s)\n", container.ID, container.Image, container.Status)
	}

	// web server
	web_server := webserver.NewWebserver(cli, log)
	web_server.Start()
}
