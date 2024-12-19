package handlers

import (
	"github.com/amirazad1/ELearning/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

type Person struct {
	Name   string  `json:"name" form:"name" binding:"required,alpha,min=3,max=6"`
	Age    int     `json:"age" form:"age" binding:"gt=0"`
	Mobile *string `json:"mobile" form:"mobile" binding:"mobile"`
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	header := c.GetHeader("my-header")
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"message":   "working!",
		"my header": header,
		"ids":       ids,
		"name":      name,
	}, true, 0))
}

func (h *HealthHandler) HealthWithParam(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"id":   id,
		"name": name,
	}, true, 0))
}

func (h *HealthHandler) HealthWithBody(c *gin.Context) {
	person := &Person{}
	err := c.ShouldBindBodyWithJSON(person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(person, false, helper.ValidationError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(person, true, 0))
}

func (h *HealthHandler) HealthWithForm(c *gin.Context) {
	person := &Person{}
	err := c.ShouldBind(person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(person, false, helper.ValidationError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(person, true, 0))
}

func (h *HealthHandler) HealthWithFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "uploads/file.txt")

	c.AbortWithStatusJSON(http.StatusBadRequest,
		helper.GenerateBaseResponseWithValidationError(file, false, helper.ValidationError, err))

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(file, true, 0))
}
