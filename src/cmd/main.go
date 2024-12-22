package main

import (
	"github.com/amirazad1/ELearning/api"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/data/cache"
	database "github.com/amirazad1/ELearning/data/db"
	"github.com/amirazad1/ELearning/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
