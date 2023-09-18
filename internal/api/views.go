package api

import (
	"log"
	"net/http"
	"strings"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// id, location, cargo, type
func GetAllContainers(containers *map[string]models.Container) func(c *gin.Context) {
	return func(c *gin.Context) {
		query := c.Request.URL.Query()

		if len(query) == 0 {
			c.HTML(http.StatusOK, "index.tmpl", containers)
			return
		}

		if len(query) > 2 {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}

		filter, exists := query["filter"]
		if !exists {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}
		field, exists := query["field"]
		if !exists {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}

		filterenContainers := make(map[string]models.Container)

		switch field[0] {
		case "id" :
			for id, container := range *containers {
				if strings.Contains(id, filter[0]) {
					filterenContainers[filter[0]] = container
				}
			}
		case "location" :
			for id, container := range *containers {
				if strings.Contains(strings.ToLower(container.CurrentLocation), strings.ToLower(filter[0])) {
					filterenContainers[id] = container
				}
			}
		case "type" :
			for id, container := range *containers {
				if strings.Contains(strings.ToLower(container.Type.Name), strings.ToLower(filter[0])) {
					filterenContainers[id] = container
				}
			}
		case "cargo" :
			for id, container := range *containers {
				if strings.Contains(strings.ToLower(container.Cargo.Name), strings.ToLower(filter[0])) {
					filterenContainers[id] = container
				}
			}
		default:
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
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
