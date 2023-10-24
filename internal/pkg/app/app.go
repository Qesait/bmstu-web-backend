package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/repository"
)

type Application struct {
	repo        *repository.Repository
	minioClient *minio.Client
	config      *config.Config
}

func (app *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.Use(ErrorHandler())

	// TODO: ошибки на русском
	// Услуги (контейнеры)
	r.GET("/containers", app.GetAllContainers)                                         // Список с поиском
	r.GET("/containers/:container_id", app.GetContainer)                               // Одна услуга
	r.DELETE("/containers/:container_id/delete", app.DeleteContainer)                  // Удаление
	r.PUT("/containers/:container_id/update", app.ChangeContainer)                     // Изменение
	r.POST("/containers/add", app.AddContainer)                                        // Добавление
	r.POST("/containers/:container_id/add_to_transportation", app.AddToTranspostation) // Добавление в заявку

	// Заявки (перевозки)
	r.GET("/transportations", app.GetAllTransportations)                                                         // Список (отфильтровать по дате формирования и статусу)
	r.GET("/transportations/:transportation_id", app.GetTranspostation)                                  // Одна заявка
	r.PUT("/transportations/:transportation_id/update", app.UpdateTransportation)                                // Изменение (добавление транспорта)
	r.DELETE("/transportations/:transportation_id/delete", app.DeleteTransportation)                             //Удаление
	r.DELETE("/transportations/:transportation_id/delete_container/:container_id", app.DeleteFromTransportation) // Изменеие (удаление услуг)
	r.PUT("/transportations/:transportation_id/user_confirm", app.UserConfirm)                                   // Сформировать создателем
	r.PUT("transportations/:transportation_id/moderator_confirm", app.ModeratorConfirm)                          // Завершить отклонить модератором

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

	app.minioClient, err = minio.New(app.config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("", "", ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			log.Println(err.Err)
		}
		lastError := c.Errors.Last()
		if lastError != nil {
			switch c.Writer.Status() {
			case http.StatusBadRequest:
				c.JSON(-1, gin.H{"error": "wrong request"})
			case http.StatusNotFound:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			case http.StatusMethodNotAllowed:
				c.JSON(-1, gin.H{"error": lastError.Error()})
			default:
				c.Status(-1)
			}
		}
	}
}
