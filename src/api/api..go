package api

import (
	"fmt"
	"github.com/amirazad1/ELearning/api/middleware"
	"github.com/amirazad1/ELearning/api/routers"
	"github.com/amirazad1/ELearning/api/validation"
	"github.com/amirazad1/ELearning/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middleware.LimitByRequest())

	RegisterRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = val.RegisterValidation("mobile", validation.IranianMobileNumberValidator)
		_ = val.RegisterValidation("password", validation.PasswordValidator)
	}
}
