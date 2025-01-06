package middleware

import (
	"github.com/amirazad1/ELearning/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, helper.CustomRecovery, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, helper.CustomRecovery, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
