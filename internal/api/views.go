package api

import (
	"log"
	"net/http"
	"strings"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAllContainers(containers *map[string]models.Container) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		log.Println(query)

		if len(query) == 0 {
			c.HTML(http.StatusOK, "index.tmpl", containers)
			return
		}

		if len(query) > 1 {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}

		containerType, exists := query["type"]
		if !exists {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}

		filterenContainers := make(map[string]models.Container)
		for id, container := range *containers {
			if strings.Contains(strings.ToLower(container.Type.Name), containerType[0]) {
				filterenContainers[id] = container
			}
		}
		c.HTML(http.StatusOK, "index.tmpl", filterenContainers)
	}
}

func GetOneContainer(containers *map[string]models.Container) func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println(c.Params)
		container, exists := (*containers)[c.Param("id")]

		if !exists {
			c.HTML(http.StatusNotFound, "error.tmpl", http.StatusNotFound)
			return
		}
		c.HTML(http.StatusOK, "item-info.tmpl", container)
	}
}
