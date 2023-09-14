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
		log.Println(query)

		if len(query) == 0 {
			c.HTML(http.StatusOK, "index.tmpl", containers)
			return
		}

		filterenContainers := make(map[string]models.Container)
		if values, exists := query["id"]; exists {
			log.Println("Look! It's id!")
			searchId := values[0]
			for id, container := range *containers {
				if strings.Contains(id, searchId) {
					filterenContainers[searchId] = container
				}
			}
			} else if values, exists := query["location"]; exists {
			log.Println("Look! It's location!")
			searchLocation := values[0]
			for id, container := range *containers {
				if strings.Contains(strings.ToLower(container.CurrentLocation), strings.ToLower(searchLocation)) {
					filterenContainers[id] = container
				}
			}
			} else if values, exists := query["type"]; exists {
			log.Println("Look! It's type!")
			searchType := values[0]
			for id, container := range *containers {
				if strings.Contains(strings.ToLower(container.Type.Name), strings.ToLower(searchType)) {
					filterenContainers[id] = container
				}
			}
			} else if values, exists := query["cargo"]; exists {
			log.Println("Look! It's cargo!")
			searchCargo := values[0]
			for id, container := range *containers {
				// Сравнение без учета регистра
				if strings.Contains(strings.ToLower(container.Cargo.Name), strings.ToLower(searchCargo)) {
					filterenContainers[id] = container
				}
			}
		} else {
			c.HTML(http.StatusBadRequest, "error.tmpl", http.StatusBadRequest)
			return
		}
		log.Println("Filtered containers:", filterenContainers)
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
