package app

import (
	_ "bmstu-web-backend/docs"
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary		Получить все контейнеры
// @Tags		Контейнеры
// @Description	Возвращает все доступные контейнеры с опциональной фильтрацией по типу
// @Produce		json
// @Param		type query string false "тип контейнера для фильтрации"
// @Success		200 {object} schemes.GetAllContainersResponse
// @Router		/api/containers [get]
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

	response := schemes.GetAllContainersResponse{DraftTransportation: nil, Containers: containers}
	if userId, exists := c.Get("userId"); exists {
		draftTransportation, err := app.repo.GetDraftTransportation(userId.(string))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if draftTransportation != nil {
			response.DraftTransportation = &draftTransportation.UUID
		}
	}
	c.JSON(http.StatusOK, response)
}

// @Summary		Получить один контейнер
// @Tags		Контейнеры
// @Description	Возвращает более подробную информацию об одном контейнере
// @Produce		json
// @Param		id path string true "id контейнера"
// @Success		200 {object} ds.Container
// @Router		/api/containers/{id} [get]
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

// @Summary		Удалить контейнер
// @Tags		Контейнеры
// @Description	Удаляет контейнер по id
// @Param		id path string true "id контейнера"
// @Success		200
// @Router		/api/containers/{id} [delete]
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
	if container.ImageURL != nil {
		if err := app.deleteImage(c, container.UUID); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	container.ImageURL = nil
	container.IsDeleted = true
	if err := app.repo.SaveContainer(container); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Добавить контейнер
// @Tags		Контейнеры
// @Description	Добавить новый контейнер
// @Accept		mpfd
// @Param		image formData file false "Изображение контейнера"
// @Param		marking formData string true "Маркировка" format:"string" maxLength:11
// @Param		type formData string true "Тип" format:"string" maxLength:50
// @Param		length formData int true "Длина" format:"int"
// @Param		height formData int true "Высота" format:"int"
// @Param		width formData int true "Ширина" format:"int"
// @Param		cargo formData string true "Груз" format:"string" maxLength:50
// @Param		weight formData int true "Вес" format:"int"
// @Success		200
// @Router		/api/containers [post]
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

	c.Status(http.StatusCreated)
}

// @Summary		Изменить котейнер
// @Tags		Контейнеры
// @Description	Изменить данные полей о контейнере
// @Accept		mpfd
// @Param		id path string true "Идентификатор контейнера" format:"uuid"
// @Param		marking formData string false "Маркировка" format:"string" maxLength:11
// @Param		type formData string false "Тип" format:"string" maxLength:50
// @Param		length formData int false "Длина" format:"int"
// @Param		height formData int false "Высота" format:"int"
// @Param		width formData int false "Ширина" format:"int"
// @Param		image formData file false "Изображение контейнера"
// @Param		cargo formData string false "Груз" format:"string" maxLength:50
// @Param		weight formData int false "Вес" format:"int"
// @Success		200
// @Router		/api/containers/{id} [put]
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

	c.Status(http.StatusOK)
}

// @Summary		Добавить в перевозку
// @Tags		Контейнеры
// @Description	Добавить выбранный контейнер в черновик перевозки
// @Param		id path string true "id контейнера"
// @Success		200
// @Router		/api/containers/{id}/add_to_transportation [post]
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
	userId := getUserId(c)
	transportation, err = app.repo.GetDraftTransportation(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		transportation, err = app.repo.CreateDraftTransportation(userId)
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

	c.Status(http.StatusOK)
}
