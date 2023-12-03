package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	var draftTransportation *ds.Transportation = nil
	if userId, exists := c.Get("userId"); exists{
		draftTransportation, err = app.repo.GetDraftTransportation(userId.(string))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	response := schemes.GetAllContainersResponse{DraftTransportation: nil, Containers: containers}
	if draftTransportation != nil {
		response.DraftTransportation = &schemes.TransportationShort{UUID: draftTransportation.UUID}
		containers, err := app.repo.GetTransportatioinComposition(draftTransportation.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		response.DraftTransportation.ContainerCount = len(containers)
	}
	c.JSON(http.StatusOK, response)
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
	if container == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("контейнер не найден"))
		return
	}
	c.JSON(http.StatusOK, container)
}

func (app *Application) DeleteContainer(c *gin.Context) {
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
	if container == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("контейнер не найден"))
		return
	}
	if err := app.deleteImage(c, container.UUID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	container.ImageURL = nil
	container.IsDeleted = true
	if err := app.repo.SaveContainer(container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) AddContainer(c *gin.Context) {
	var request schemes.AddContainerRequest
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	container := ds.Container(request.Container)
	if err := app.repo.AddContainer(&container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if request.Image != nil {
		imageURL, err := app.uploadImage(c, request.Image, container.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		container.ImageURL = imageURL
	}
	if err := app.repo.SaveContainer(&container); err != nil {
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
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if container == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("контейнер не найден"))
		return
	}
	if request.Marking != nil {
		container.Marking = *request.Marking
	}
	if request.Type != nil {
		container.Type = *request.Type
	}
	if request.Length != nil {
		container.Length = *request.Length
	}
	if request.Height != nil {
		container.Height = *request.Height
	}
	if request.Width != nil {
		container.Width = *request.Width
	}
	if request.Cargo != nil {
		container.Cargo = *request.Cargo
	}
	if request.Weight != nil {
		container.Weight = *request.Weight
	}
	if request.Image != nil {
		if container.ImageURL != nil {
			if err := app.deleteImage(c, container.UUID); err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
		imageURL, err := app.uploadImage(c, request.Image, container.UUID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		container.ImageURL = imageURL
	}

	if err := app.repo.SaveContainer(container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, container)
}

func (app *Application) AddToTranspostation(c *gin.Context) {
	var request schemes.AddToTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var err error

	// Проверить существует ли контейнер
	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if container == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("контейнер не найден"))
		return
	}

	// Получить черновую заявку
	var transportation *ds.Transportation
	userId, _ := c.Get("userId")
	transportation, err = app.repo.GetDraftTransportation(userId.(string))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		transportation, err = app.repo.CreateDraftTransportation(userId.(string))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Создать связь меджду перевозкой и контейнером
	if err = app.repo.AddToTransportation(transportation.UUID, request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Вернуть список всех контейнеров в перевозке
	var containers []ds.Container
	containers, err = app.repo.GetTransportatioinComposition(transportation.UUID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}
