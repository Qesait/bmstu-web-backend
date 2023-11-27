package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (app *Application) GetAllTransportations(c *gin.Context) {
	var request schemes.GetAllTransportationsRequst
	if err := c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportations, err := app.repo.GetAllTransportations(request.FormationDateStart, request.FormationDateEnd, request.Status)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	outputTransportations := make([]schemes.TransportationOutput, len(transportations))
	for i, transportation := range transportations {
		outputTransportations[i] = schemes.ConvertTransportation(&transportation)
	}
	c.JSON(http.StatusOK, schemes.AllTransportationsResponse{Transportations: outputTransportations})
}

func (app *Application) GetTranspostation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: schemes.ConvertTransportation(transportation), Containers: containers})
}

func (app *Application) UpdateTransportation(c *gin.Context) {
	var request schemes.UpdateTransportationRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	transportation.Transport = request.Transport
	if app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.UpdateTransportationResponse{Transportation: schemes.ConvertTransportation(transportation)})
}

func (app *Application) DeleteTransportation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	transportation.Status = ds.DELETED

	if err := app.repo.SaveTransportation(transportation); err != nil {
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
	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	if transportation.Status != ds.DRAFT {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя редактировать перевозку со статусом: %s", transportation.Status))
		return
	}

	if err := app.repo.DeleteFromTransportation(request.TransportationId, request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) UserConfirm(c *gin.Context) {
	var request schemes.UserConfirmRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	if transportation.Status != ds.DRAFT {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя сформировать перевозку со статусом %s", transportation.Status))
		return
	}

	if request.Confirm {
		transportation.Status = ds.FORMED
		now := time.Now()
		transportation.FormationDate = &now
	} else {
		transportation.Status = ds.DELETED
	}

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) ModeratorConfirm(c *gin.Context) {
	var request schemes.ModeratorConfirmRequest
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	if transportation.Status != ds.FORMED {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить статус с \"%s\" на \"%s\"", transportation.Status, ds.FORMED))
		return
	}

	if request.Confirm {
		transportation.Status = ds.COMPELTED
		now := time.Now()
		transportation.CompletionDate = &now
	} else {
		transportation.Status = ds.REJECTED
	}
	transportation.ModeratorId = app.getModerator()

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
