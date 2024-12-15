package routers

import (
	"github.com/amirazad1/ELearning/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup){
	handler:=handlers.NewHealthHandler()

	r.GET("/",handler.Health)	
}