package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/schemes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) GetContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		log.Println("can't get product by id", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if container == nil {
		log.Println("no container with id", err)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, schemes.GetContainerResponse{Container: *container})
}

func (app *Application) GetAllContainers(c *gin.Context) {
	var request schemes.GetAllContainersRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Println("can't parse request query params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	containers, err := app.repo.GetContainersByType(request.ContainerType)
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) DeleteContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := app.repo.DeleteContainer(request.ContainerId); err != nil {
		log.Println("can't delete container from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetContainersByType("")
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) AddToTranspostation(c *gin.Context) {
	var request schemes.AddToTransportationRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("can't parse requst body", err)
		c.Status(http.StatusBadRequest)
		return
	}
	var err error

	var transportation *ds.Transportation
	transportation, err = app.repo.GetEditableTransportation()
	if err != nil {
		log.Println("can't get transportation from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if err = app.repo.AddToTransportation(transportation.UUID, request.ContainerId); err != nil {
		log.Println("can't add container to transportation in db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	var containers []ds.Container
	containers, err = app.repo.GetTransportatioinComposition(transportation.UUID)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) TranspostationComposition(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.Transpostationid)
	if err != nil {
		log.Printf("can't get transportation by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.Transpostationid)
	if err != nil {
		log.Printf("can't get transportation composition by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) UpdateTransportation(c *gin.Context) {
	var request schemes.UpdateTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("can't parse request body", err)
		c.Status(http.StatusBadRequest)
		return
	}
	transportation, err := app.repo.GetTransportationById(request.TransportationId)
	if err != nil {
		log.Printf("can't get transportation by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	transportation.Transport = request.Transport
	if app.repo.SaveTransportation(transportation); err != nil {
		log.Printf("can't save transportation %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) DeleteTransportation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := app.repo.DeleteTransportation(request.Transpostationid)
	if err != nil {
		log.Println("can't delete transportation from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) DeleteFromTransportation(c *gin.Context) {
	var request schemes.DeleteFromTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := app.repo.DeleteFromTransportation(request.Transpostationid, request.ContainerId); err != nil {
		log.Println("can't delete container from transportation in db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
