package api

import (
	"net/http"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllContainers(containers map[string]models.Container) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", containers)
	}
}

func GetOneContainer(containers map[string]models.Container) func(c *gin.Context) {
	return func(c *gin.Context) {
		container, exists := containers[c.Param("id")]

		if !exists {
			c.HTML(http.StatusNotFound, "error.tmpl", http.StatusNotFound)
			return
		}
		c.HTML(http.StatusOK, "item-info.tmpl", container)
	}
}
