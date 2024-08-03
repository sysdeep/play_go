package main

import (
	"hdu/internal/registry/registry_client"
	"hdu/internal/utils"
)

func main() {
	client := registry_client.NewHTTPRegistryClient("https://localhost:5000")
	catalog := client.GetCatalog(1)
	utils.PrintAsJson(catalog)
}
