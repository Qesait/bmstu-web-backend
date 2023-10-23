package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *Application) getUser() string {
	return "5f58c307-a3f2-4b13-b888-c80ad08d5ed3"
}

func (app *Application) GetAllContainerTypes(c *gin.Context) {
	types, err := app.repo.GetAllContainerTypes()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.AllContainerTypesResponse{ContainerTypes: types})
}

func (app *Application) GetContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.GetContainerResponse{Container: *container})
}

func (app *Application) GetAllContainers(c *gin.Context) {
	var request schemes.GetAllContainersRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	containers, err := app.repo.GetContainersByType(request.ContainerType)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) DeleteContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := app.repo.DeleteContainer(request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	containers, err := app.repo.GetContainersByType("")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) AddContainer(c *gin.Context) {
	container := &ds.Container{}
	if err := c.ShouldBindJSON(container); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := app.repo.AddContainer(container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) ChangeContainer(c *gin.Context) {
	var request schemes.ChangeContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if request.TypeId != nil {
		container.TypeId = *request.TypeId
		containerType, err := app.repo.GetContainerType(*request.TypeId)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		container.ContainerType = *containerType
	}
	if request.ImageURL != nil {
		container.ImageURL = *request.ImageURL
	}
	if request.PurchaseDate != nil {
		container.PurchaseDate = *request.PurchaseDate
	}
	if request.Cargo != nil {
		container.Cargo = *request.Cargo
	}
	if request.Weight != nil {
		container.Weight = *request.Weight
	}
	if request.Marking != nil {
		container.Marking = *request.Marking
	}

	if err := app.repo.SaveContainer(container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) AddToTranspostation(c *gin.Context) {
	var request schemes.AddToTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var err error

	var transportation *ds.Transportation
	transportation, err = app.repo.GetEditableTransportation(app.getUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = app.repo.AddToTransportation(transportation.UUID, request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var containers []ds.Container
	containers, err = app.repo.GetTransportatioinComposition(transportation.UUID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) GetAllTransportations(c *gin.Context) {
	var request schemes.GetAllTransportationsRequst
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportations, err := app.repo.GetAllTransportations(request.FormationDate, request.Status)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println(transportations)
	c.JSON(http.StatusOK, schemes.AllTransportationsResponse{Transportations: transportations})
}

func (app *Application) TranspostationComposition(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) UpdateTransportation(c *gin.Context) {
	var request schemes.UpdateTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	transportation.Transport = request.Transport
	if app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) DeleteTransportation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := app.repo.UpdateTransportationStatus(request.TransportationId, app.getUser(), ds.DELETED)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) DeleteFromTransportation(c *gin.Context) {
	var request schemes.DeleteFromTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := app.repo.DeleteFromTransportation(request.TransportationId, request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) GetContainerType(c *gin.Context) {
	var request schemes.TypeRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	containerType, err := app.repo.GetContainerType(request.TypeId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.GetTypeResponse{ContainerType: *containerType})
}

func (app *Application) UserConfirm(c *gin.Context) {
	var request schemes.UserConfirmRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := app.repo.UpdateTransportationStatus(request.TransportationId, app.getUser(), "сформирован")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) ModeratorConfirm(c *gin.Context) {
	var request schemes.ModeratorConfirmRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getUser())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if transportation.Status != ds.FORMED {
		c.AbortWithError(http.StatusMethodNotAllowed, err)
		return
	}

	transportation.Status = request.Status
	if request.Status == ds.COMPELTED {
		now := time.Now()
		transportation.CompletionDate = &now
	}

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
