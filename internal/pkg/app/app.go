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

	// Containers API
	r.GET("/containers", app.GetContainers)
	r.GET("/containers/:id", app.GetContainer)
	r.DELETE("/containers/:id/delete", app.DeleteContainer)
	// Transportation API
	r.POST("/transportation", app.AddToTranspostation)
	r.DELETE("/transportation/:transportation_id/:container_id/delete", app.DeleteFromTransportation)
	r.PUT("/transportation/:transportation_id/put", app.UpdateTransportation)
	r.DELETE("/transportation/:transportation_id/delete", app.DeleteTransportation)
	
	// TODO: убрать
	r.GET("/transportation/:transportation_id", app.TranspostationComposition)

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
