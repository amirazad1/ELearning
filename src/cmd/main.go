package main

import (
	"github.com/amirazad1/ELearning/api"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/data/cache"
)

func main() {
	cfg := config.GetConfig()
	_ = cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
