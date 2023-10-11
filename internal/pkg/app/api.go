package app

import (
	"bmstu-web-backend/internal/app/ds"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func (app *Application) GetContainer(c *gin.Context) {
	id := c.Param("id")

	container, err := app.repo.GetContainerByID(id)
	if err != nil {
		log.Printf("can't get product by id %v", err)
		c.Error(err)
		return
	}

	if len(container.Cargo) == 0 {
		log.Println("empty")
	}

	c.JSON(http.StatusOK, *container)
}

func (app *Application) GetContainers(c *gin.Context) {
	containerType := c.Query("type")

	containers, err := app.repo.GetContainersByType(containerType)
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}

func (app *Application) DecommissionContainer(c *gin.Context) {
	id := c.PostForm("delete")

	app.repo.DecommissionContainer(id)

	containers, err := app.repo.GetContainersByType("")
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}

func (app *Application) AddToTranspostation(c *gin.Context) {
	containerId := uuid.MustParse(c.PostForm("container_id"))
	log.Println(containerId)
	var err error
	var transportation *ds.Transportation
	
	transportation, err = app.repo.GetEditableTransportation()
	if err != nil {
		log.Println("can't get transportation from db", err)
		c.Error(err)
		return
	}
	if transportation == nil {
		transportation, err = app.repo.CreateTransportation()
		if err != nil {
			log.Println("can't create transportation in db", err)
			c.Error(err)
			return
		}
	}
	log.Println(*transportation)
	err = app.repo.AddContainerToTransportation(transportation.UUID, containerId)
	if err != nil {
		log.Println("can't add container to transportation in db", err)
		c.Error(err)
		return
	}

	var containers []ds.TransportationComposition
	containers, err = app.repo.GetTransportatioinComposition(transportation)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}