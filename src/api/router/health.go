package router

import (
	"github.com/amirazad1/ELearning/api/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
	r.GET("/:id/:name", handler.HealthWithParam)
	r.POST("/", handler.HealthWithBody)
	r.POST("/form", handler.HealthWithForm)
	r.POST("/file", handler.HealthWithFile)
}
