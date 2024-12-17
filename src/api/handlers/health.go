package handlers

import (
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

	c.JSON(http.StatusOK, gin.H{
		"message":   "working!",
		"my header": header,
		"ids":       ids,
		"name":      name,
	})
}

func (h *HealthHandler) HealthWithParam(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func (h *HealthHandler) HealthWithBody(c *gin.Context) {
	person := &Person{}
	err := c.ShouldBindBodyWithJSON(person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func (h *HealthHandler) HealthWithForm(c *gin.Context) {
	person := &Person{}
	err := c.ShouldBind(person)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"person": person,
	})
}

func (h *HealthHandler) HealthWithFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "uploads/file.txt")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file": file.Filename,
		"size": file.Size,
	})
}
