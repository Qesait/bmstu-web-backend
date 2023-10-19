package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
}

func (app *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	// Типы контейнеров
	r.GET("/container_types", app.GetAllContainerTypes) // Список типов
	r.GET("/container_types/:type_id", app.GetContainerType) // один тип
	// Услуги (контейнеры)
	r.GET("/containers", app.GetAllContainers)                        // Список с поиском
	r.GET("/containers/:container_id", app.GetContainer)              // Одна услуга
	r.DELETE("/containers/:container_id/delete", app.DeleteContainer) // Удаление
	r.PUT("/containers/:container_id/put", app.ChangeContainer)       // Изменение
	r.POST("/containers", app.AddContainer)                           // Добавление
	r.POST("/transportations", app.AddToTranspostation)               // Добавление в заявку

	// Заявки (перевозки)
	r.GET("/transportations", app.GetAllTransportations)                                               // Список (отфильтровать по дате формирования и статусу)
	r.GET("/transportations/:transportation_id", app.TranspostationComposition)                        // Одна заявка
	r.DELETE("/transportations/:transportation_id/:container_id/delete", app.DeleteFromTransportation) // Изменеие (удаление услуг)
	r.PUT("/transportations/:transportation_id/put", app.UpdateTransportation)                         // Изменение (добавление транспорта)
	r.DELETE("/transportations/:transportation_id/delete", app.DeleteTransportation)                   //Удаление
	// Сформировать создателем
	// Завершить отклонить модератором

	r.Static("/image", "./static/image")
	r.Static("/css", "./static/css")

	r.Run(fmt.Sprintf("%s:%d", app.config.ServiceHost, app.config.ServicePort))

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	return &app, nil
}
