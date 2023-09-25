package api

import (
	db "bmstu-web-backend/internal/database"
	"bmstu-web-backend/internal/models"
	"log"

	"github.com/gin-gonic/gin"
)

type server struct {
	db     map[string]models.Container
	engine *gin.Engine
}

func StartServer() {
	log.Println("Server start up")
	s := server{
		db:     db.SetupDB(),
		engine: gin.Default(),
	}

	s.engine.LoadHTMLGlob("templates/*")

	s.engine.GET("/containers", GetAllContainers(&s.db))
	s.engine.GET("/containers/:id", GetOneContainer(&s.db))

	s.engine.Static("/image", "./static/image")
	s.engine.Static("/css", "./static/css")

	s.engine.Run()
	log.Println("Server down")
}
