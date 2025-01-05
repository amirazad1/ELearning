package handler

import (
	"github.com/amirazad1/ELearning/api/dto"
	"github.com/amirazad1/ELearning/api/helper"
	"github.com/amirazad1/ELearning/config"
	"github.com/amirazad1/ELearning/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersHandler struct {
	userService *services.UserService
	otpService  *services.OtpService
}

func NewUserHandler(cfg *config.Config) *UsersHandler {
	userService := services.NewUserService(cfg)
	otpService := services.NewOtpService(cfg)
	return &UsersHandler{
		userService: userService,
		otpService:  otpService,
	}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err = h.otpService.SendOtp(req.MobileNumber)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	// TODO: Call internal SMS service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}
