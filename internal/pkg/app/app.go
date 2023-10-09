package app

import (
	"log"

	"github.com/gin-gonic/gin"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
}

func (app *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/containers/:id", app.GetContainer())
	r.GET("/containers", app.GetContainers())
	r.POST("/containers", app.DecommissionContainer())

	r.LoadHTMLGlob("templates/*")

	r.Static("/image", "./static/image")
	r.Static("/css", "./static/css")

	r.Run()

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	return &app, nil
}
