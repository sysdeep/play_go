package main

import (
	"context"
	"fmt"
	"hdu/internal/webserver"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	fmt.Println("start")

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
	web_server := webserver.NewWebserver(cli)
	web_server.Start()
}
