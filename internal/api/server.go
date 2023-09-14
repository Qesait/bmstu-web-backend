package api

import (
	"log"

	"bmstu-web-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")

	containers := map[string]models.Container{
		"AAAU1234560": {
			Id: "AAAU1234560",
			Type:         models.DRY_CUBE_20,
			ImageURL:     "http://localhost:8080/image/0.jpeg",
			Cargo: models.Cargo{
				Name:   "Бумага",
				Weight: 20000,
			},
			CurrentLocation: "Архангельск",
		},
		"BBBU6543210": {
			Id: "BBBU6543210",
			Type:         models.HIGH_CUBE_20,
			ImageURL:     "http://localhost:8080/image/1.jpg",
			Cargo: models.Cargo{
				Name:   "Зерно",
				Weight: 15000,
			},
			CurrentLocation: "Калининград",
		},
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/containers", GetAllContainers(&containers))
	r.GET("/containers/:id", GetOneContainer(&containers))

	r.Static("/image", "./resources")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}
