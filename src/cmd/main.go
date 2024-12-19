package main

import (
	"github.com/amirazad1/ELearning/api"
	"github.com/amirazad1/ELearning/config"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
}
