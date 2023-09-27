package api

import (
	db "bmstu-web-backend/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	log.Println("Server start up")
	db := db.SetupDB()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/containers", GetAllContainers(&db))
	r.GET("/containers/:id", GetOneContainer(&db))

	r.Static("/image", "./static/image")
	r.Static("/css", "./static/css")

	r.Run()
	log.Println("Server down")
}
