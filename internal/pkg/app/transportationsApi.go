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
// @Param		transportation_id path string true "id перевозки"
// @Success		200 {object} schemes.TransportationResponse
// @Router		/api/transportations/{transportation_id} [get]
func (app *Application) GetTranspostation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	transportation, err := app.repo.GetTransportationById(request.TransportationId, userId)
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
// @Description	Позволяет изменить транспорт перевозки и возвращает обновлённые данные
// @Access		json
// @Produce		json
// @Param		transportation_id path string true "id перевозки"
// @Param		transport body SwaggerUpdateTransportationRequest true "Транспорт"
// @Success		200 {object} schemes.UpdateTransportationResponse
// @Router		/api/transportations/{transportation_id} [put]
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

	userId := getUserId(c)
	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, userId)
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

// @Summary		Удалить перевозку
// @Tags		Перевозки
// @Description	Удаляет первозку по id
// @Param		transportation_id path string true "id перевозки"
// @Success		200
// @Router		/api/transportations/{transportation_id} [delete]
func (app *Application) DeleteTransportation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	transportation, err := app.repo.GetTransportationById(request.TransportationId, userId)
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

// @Summary		Удалить контейнер из перевозки
// @Tags		Перевозки
// @Description	Удалить контейнер из перевозки
// @Produce		json
// @Param		transportation_id path string true "id перевозки"
// @Param		container_id path string true "id контейнера"
// @Success		200 {object} schemes.AllContainersResponse
// @Router		/api/transportations/{transportation_id}/delete_container/{container_id} [delete]
func (app *Application) DeleteFromTransportation(c *gin.Context) {
	var request schemes.DeleteFromTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId := getUserId(c)
	transportation, err := app.repo.GetTransportationById(request.TransportationId, userId)
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

// @Summary		Сформировать перевозку
// @Tags		Перевозки
// @Description	Сформировать или удалить перевозку перевозку пользователем
// @Produce		json
// @Param		confirm body boolean true "подтвердить"
// @Success		200
// @Router		/api/transportations/user_confirm [put]
func (app *Application) UserConfirm(c *gin.Context) {
	var request schemes.UserConfirmRequest
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

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

// @Summary		Подтвердить перевозку
// @Tags		Перевозки
// @Description	Подтвердить или отменить перевозку модератором
// @Produce		json
// @Param		transportation_id path string true "id перевозки"
// @Param		confirm body boolean true "подтвердить"
// @Success		200
// @Router		/api/transportations/{transportation_id}/moderator_confirm [put]
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
	transportation, err := app.repo.GetTransportationById(request.URI.TransportationId, userId)
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
	transportation.ModeratorId = &userId

	if err := app.repo.SaveTransportation(transportation); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
