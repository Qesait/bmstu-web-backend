package api

import (
	"log"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")

	containers := map[string]models.Container{
		"ys7E4rnOxo": {
			SerialNumber: "ys7E4rnOxo",
			Type:         "20 футов Dry Cube",
			ImageURL:     "http://localhost:8080/image/0.jpeg",
			Dimentions:   models.DRY_CUBE_20,
			Cargo: models.Cargo{
				Name:   "Бумага",
				Weight: 20000,
			},
		},
		"2TdWlrFGe5": {
			SerialNumber: "2TdWlrFGe5",
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

	r.GET("/containers", GetAllContainers(containers))
	r.GET("/containers/:id", GetOneContainer(containers))

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
