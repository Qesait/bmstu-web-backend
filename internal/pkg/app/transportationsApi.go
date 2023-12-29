package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/role"
	"bmstu-web-backend/internal/app/schemes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary		Получить все перевозки
// @Tags		Перевозки
// @Description	Возвращает все перевозки с фильтрацией по статусу и дате формирования
// @Produce		json
// @Param		status query string false "статус перевозки"
// @Param		formation_date_start query string false "начальная дата формирования"
// @Param		formation_date_end query string false "конечная дата формирвания"
// @Success		200 {object} schemes.AllTransportationsResponse
// @Router		/api/transportations [get]
func (app *Application) GetAllTransportations(c *gin.Context) {
	var request schemes.GetAllTransportationsRequst
	var err error
	if err = c.ShouldBindQuery(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	userRole := getUserRole(c)
	fmt.Println(userId, userRole)
	var transportations []ds.Transportation
	if userRole == role.Customer {
		transportations, err = app.repo.GetAllTransportations(&userId, request.FormationDateStart, request.FormationDateEnd, request.Status)
	} else {
		transportations, err = app.repo.GetAllTransportations(nil, request.FormationDateStart, request.FormationDateEnd, request.Status)
	}
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

// @Summary		Получить одну перевозку
// @Tags		Перевозки
// @Description	Возвращает подробную информацию о перевозке и её составе
// @Produce		json
// @Param		id path string true "id перевозки"
// @Success		200 {object} schemes.TransportationResponse
// @Router		/api/transportations/{id} [get]
func (app *Application) GetTranspostation(c *gin.Context) {
	var request schemes.TranspostationRequest
	var err error
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	userRole := getUserRole(c)
	var transportation *ds.Transportation
	if userRole == role.Moderator {
		transportation, err = app.repo.GetTransportationById(request.TransportationId, nil)
	} else {
		transportation, err = app.repo.GetTransportationById(request.TransportationId, &userId)
	}
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

type SwaggerUpdateTransportationRequest struct {
	Transport string `json:"transport"`
}

// @Summary		Указать транспорт перевозки
// @Tags		Перевозки
// @Description	Позволяет изменить транспорт черновой перевозки и возвращает обновлённые данные
// @Access		json
// @Produce		json
// @Param		transport body SwaggerUpdateTransportationRequest true "Транспорт"
// @Success		200
// @Router		/api/transportations [put]
func (app *Application) UpdateTransportation(c *gin.Context) {
	var request schemes.UpdateTransportationRequest
	var err error
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	// Добавить транспорт
	transportation.Transport = &request.Transport
	if app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Удалить черновую первозку перевозку
// @Tags		Перевозки
// @Description	Удаляет чернвоую перевозку первозку
// @Success		200
// @Router		/api/transportations [delete]
func (app *Application) DeleteTransportation(c *gin.Context) {
	var err error

	// Получить черновую заявку
	var transportation *ds.Transportation
	userId := getUserId(c)
	transportation, err = app.repo.GetDraftTransportation(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	transportation.Status = ds.StatusDeleted

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}

// @Summary		Удалить контейнер из чернвоой перевозки
// @Tags		Перевозки
// @Description	Удалить контейнер из черновой перевозки
// @Produce		json
// @Param		id path string true "id контейнера"
// @Success		200
// @Router		/api/transportations/delete_container/{id} [delete]
func (app *Application) DeleteFromTransportation(c *gin.Context) {
	var request schemes.DeleteFromTransportationRequest
	var err error
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
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
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	if err := app.repo.DeleteFromTransportation(transportation.UUID, request.ContainerId); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Сформировать перевозку
// @Tags		Перевозки
// @Description	Сформировать перевозку перевозку пользователем
// @Success		200
// @Router		/api/transportations/user_confirm [put]
func (app *Application) UserConfirm(c *gin.Context) {
	userId := getUserId(c)
	transportation, err := app.repo.GetDraftTransportation(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}

	if err := deliveryRequest(transportation.UUID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(`delivery service is unavailable: {%s}`, err))
		return
	}

	deliveryStatus := ds.DeliveryStarted
	transportation.DeliveryStatus = &deliveryStatus
	transportation.Status = ds.StatusFormed
	now := time.Now()
	transportation.FormationDate = &now

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// @Summary		Подтвердить перевозку
// @Tags		Перевозки
// @Description	Подтвердить или отменить перевозку модератором
// @Param		id path string true "id перевозки"
// @Param		confirm body boolean true "подтвердить"
// @Success		200 {object} schemes.TransportationOutput
// @Router		/api/transportations/{id}/moderator_confirm [put]
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

	userId := getUserId(c)
	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	if transportation.Status != ds.StatusFormed {
		c.AbortWithError(http.StatusMethodNotAllowed, fmt.Errorf("нельзя изменить статус с \"%s\" на \"%s\"", transportation.Status, ds.StatusFormed))
		return
	}

	if *request.Confirm {
		transportation.Status = ds.StatusCompleted
		now := time.Now()
		transportation.CompletionDate = &now
	} else {
		transportation.Status = ds.StatusRejected
	}

	moderator, err := app.repo.GetUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	transportation.ModeratorId = &userId
	transportation.Moderator = moderator

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schemes.ConvertTransportation(transportation))
}

func (app *Application) Delivery(c *gin.Context) {
	var request schemes.DeliveryReq
	if err := c.ShouldBindUri(&request.URI); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if request.Token != app.config.Token {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if transportation == nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("перевозка не найдена"))
		return
	}
	// if transportation.Status != ds.StatusFormed || *transportation.DeliveryStatus != ds.DeliveryStarted {
	// 	c.AbortWithStatus(http.StatusMethodNotAllowed)
	// 	return
	// }

	var deliveryStatus string
	if *request.DeliveryStatus {
		deliveryStatus = ds.DeliveryCompleted
	} else {
		deliveryStatus = ds.DeliveryFailed
	}
	transportation.DeliveryStatus = &deliveryStatus

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
