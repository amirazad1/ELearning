package api

import (	
	"github.com/amirazad1/ELearning/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	r:=gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	v1:=r.Group("/api/v1")
	{
		health:=v1.Group("/health")
		routers.Health(health)
	}

	r.Run(":5005")
}