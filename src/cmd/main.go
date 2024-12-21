package main

import (
	"github.com/amirazad1/ELearning/api"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/data/cache"
	database "github.com/amirazad1/ELearning/data/db"
	"log"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}
