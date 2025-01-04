package router

import (
	"github.com/amirazad1/ELearning/api/handler"
	"github.com/amirazad1/ELearning/api/middleware"
	"github.com/amirazad1/ELearning/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)

	router.POST("/send-otp", middleware.OtpLimiter(cfg), h.SendOtp)
}
