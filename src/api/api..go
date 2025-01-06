package api

import (
	"fmt"
	"github.com/amirazad1/ELearning/api/middleware"
	"github.com/amirazad1/ELearning/api/router"
	"github.com/amirazad1/ELearning/api/validation"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/docs"
	"github.com/amirazad1/ELearning/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterValidators()

	r.Use(middleware.DefaultStructuredLogger(cfg))
	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.CustomRecovery(middleware.ErrorHandler), middleware.LimitByRequest())

	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)

	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Started", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		users := v1.Group("/users")
		router.Health(health)
		router.User(users, cfg)
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = val.RegisterValidation("mobile", validation.IranianMobileNumberValidator)
		_ = val.RegisterValidation("password", validation.PasswordValidator)
	}
}
