package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/repository"
	"bmstu-web-backend/internal/app/role"
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

	// Услуги (контейнеры)
	r.GET("/api/containers", app.GetAllContainers)                                         // Список с поиском
	r.GET("/api/containers/:container_id", app.GetContainer)                               // Одна услуга
	r.DELETE("/api/containers/:container_id", app.DeleteContainer)                         // Удаление
	r.PUT("/api/containers/:container_id", app.ChangeContainer)                            // Изменение
	r.POST("/api/containers", app.AddContainer)                                            // Добавление
	r.POST("/api/containers/:container_id/add_to_transportation", app.AddToTranspostation) // Добавление в заявку

	// Заявки (перевозки)
	r.GET("/api/transportations", app.GetAllTransportations)                                                         // Список (отфильтровать по дате формирования и статусу)
	r.GET("/api/transportations/:transportation_id", app.GetTranspostation)                                          // Одна заявка
	r.PUT("/api/transportations/:transportation_id/update", app.UpdateTransportation)                                // Изменение (добавление транспорта)
	r.DELETE("/api/transportations/:transportation_id", app.DeleteTransportation)                                    //Удаление
	r.DELETE("/api/transportations/:transportation_id/delete_container/:container_id", app.DeleteFromTransportation) // Изменеие (удаление услуг)
	r.PUT("/api/transportations/:transportation_id/user_confirm", app.UserConfirm)                                   // Сформировать создателем
	r.PUT("/api/transportations/:transportation_id/moderator_confirm", app.ModeratorConfirm)                         // Завершить отклонить модератором

	r.POST("/api/sign_up", app.Register)
	r.POST("/api/login", app.Login) // там где мы ранее уже заводили эндпоинты
	// никто не имеет доступа
	r.GET("/api/ping", app.WithAuthCheck(role.Moderator), app.Ping)
	// или ниженаписанное значит что доступ имеют менеджер и админ
	// r.Use(a.WithAuthCheck(role.Manager, role.Admin)).GET("/ping", a.Ping)


	r.Static("/image", "./static/image")
	r.Static("/css", "./static/css")

	r.Run(fmt.Sprintf("%s:%d", app.config.ServiceHost, app.config.ServicePort))

	log.Println("Server down")
}

func New() (*Application, error) {
	var err error
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	app.repo, err = repository.New(dsn.FromEnv())
	if err != nil {
		return nil, err
	}

	app.minioClient, err = minio.New(app.config.Minio.Endpoint, &minio.Options{
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
