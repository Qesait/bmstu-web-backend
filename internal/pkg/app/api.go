package app

import (
	"bmstu-web-backend/internal/app/ds"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *Application) GetContainer(c *gin.Context) {
	id := c.Param("id")

	container, err := app.repo.GetContainerByID(id)
	if err != nil {
		log.Printf("can't get product by id %v", err)
		c.Error(err)
		return
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
	var err error
	var transportation *ds.Transportation
	containerId, err := uuid.Parse(c.PostForm("container_id"))
	if err != nil {
		log.Println("can't parse uuid", err)
		c.Error(err)
		return
	}

	transportation, err = app.repo.GetEditableTransportation()
	if err != nil {
		log.Println("can't get transportation from db", err)
		c.Error(err)
		return
	}

	err = app.repo.AddContainerToTransportation(transportation.UUID, containerId)
	if err != nil {
		log.Println("can't add container to transportation in db", err)
		c.Error(err)
		return
	}

	var containers []ds.TransportationComposition
	containers, err = app.repo.GetTransportatioinComposition(transportation.UUID)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}

func (app *Application) TranspostationComposition(c *gin.Context) {
	transportationId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("can't parse uuid", err)
		c.Error(err)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(transportationId)
	if err != nil {
		log.Printf("can't get product by id %v", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}

func (app *Application) UpdateTransportation(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println("can't parse uuid", err)
		c.Error(err)
		return
	}
	transportVehicle := c.PostForm("transport_vehicle")

	app.repo.AddTransportVehicle(id, transportVehicle)

	var containers []ds.TransportationComposition
	containers, err = app.repo.GetTransportatioinComposition(id)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, containers)
}