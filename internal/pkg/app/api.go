package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (app *Application) uploadImage(c *gin.Context, image *multipart.FileHeader, UUID string) (*string, error) {
	src, err := image.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	extension := filepath.Ext(image.Filename)
	if extension != ".jpg" && extension != ".jpeg" {
		return nil, fmt.Errorf("разрешены только jpeg изображения")
	}
	imageName := UUID + extension

	_, err = app.minioClient.PutObject(c, app.config.BucketName, imageName, src, image.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return nil, err
	}
	imageURL := fmt.Sprintf("%s/%s/%s", app.config.MinioEndpoint, app.config.BucketName, imageName)
	return &imageURL, nil
}

func (app *Application) getCustomer() string {
	return "5f58c307-a3f2-4b13-b888-c80ad08d5ed3"
}

func (app *Application) getModerator() *string {
	moderaorId := "796c70e1-5f27-4433-a415-95e7272effa5"
	return &moderaorId
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

	draftTransportation, err := app.repo.GetDraftTransportation(app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
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
	transportation, err = app.repo.GetDraftTransportation(app.getCustomer())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		transportation, err = app.repo.CreateDraftTransportation(app.getCustomer())
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
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя сформировать перевозку со статусом %s", transportation.Status))
		return
	}
	transportation.Status = ds.FORMED
	now := time.Now()
	transportation.FormationDate = &now

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

	if request.Status != ds.COMPELTED && request.Status != ds.REJECTED {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("status %s not allowed", request.Status))
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
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить статус с \"%s\" на \"%s\"", transportation.Status, request.Status))
		return
	}
	transportation.Status = request.Status
	transportation.ModeratorId = app.getModerator()
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
