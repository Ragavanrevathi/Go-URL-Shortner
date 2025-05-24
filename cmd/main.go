package main

import (
	"shorten-url/internal/routes"
	"shorten-url/pkg/config"
)

func main() {
	config.Init()
	routes.SetUpRoutesForURLShortner()
}
