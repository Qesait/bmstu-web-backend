package api

import (
	"log"
	"net/http"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
	"strconv"
)

func StartServer() {
	log.Println("Server start up")

	containers := []models.Container{
		{
			SerialNumber: 0,
			Type:         "20 футов Dry Cube",
			ImageURL:     "http://localhost:8080/image/0.jpeg",
			Dimentions:   models.DRY_CUBE_20,
			Cargo: models.Cargo{
				Name:   "Бумага",
				Weight: 20000,
			},
		},
		{
			SerialNumber: 1,
			Type:         "20 футов High Cube",
			ImageURL:     "http://localhost:8080/image/1.jpg",
			Dimentions:   models.HIGH_CUBE_20,
			Cargo: models.Cargo{
				Name:   "Зерно",
				Weight: 15000,
			},
		},
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/containers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", containers)
	})
	r.GET("/containers/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		c.HTML(http.StatusOK, "item-info.tmpl", containers[id])
	})

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
