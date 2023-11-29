package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/redis"
	"bmstu-web-backend/internal/app/repository"
	"bmstu-web-backend/internal/app/role"
)

type Application struct {
	repo        *repository.Repository
	minioClient *minio.Client
	config      *config.Config
	redisClient *redis.Client
}

func (app *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.Use(ErrorHandler())

	// Услуги (контейнеры)
	containers := r.Group("/api/containers")
	{
		containers.GET("/", app.GetAllContainers)                                        // Список с поиском
		containers.GET("/:container_id", app.GetContainer)                               // Одна услуга
		containers.DELETE("/:container_id", app.DeleteContainer)                         // Удаление
		containers.PUT("/:container_id", app.ChangeContainer)                            // Изменение
		containers.POST("/", app.AddContainer)                                           // Добавление
		containers.POST("/:container_id/add_to_transportation", app.AddToTranspostation) // Добавление в заявку
	}

	// Заявки (перевозки)
	transportations := r.Group("/api/transportations")
	{
		transportations.GET("/", app.GetAllTransportations)                                                        // Список (отфильтровать по дате формирования и статусу)
		transportations.GET("/:transportation_id", app.GetTranspostation)                                          // Одна заявка
		transportations.PUT("/:transportation_id/update", app.UpdateTransportation)                                // Изменение (добавление транспорта)
		transportations.DELETE("/:transportation_id", app.DeleteTransportation)                                    //Удаление
		transportations.DELETE("/:transportation_id/delete_container/:container_id", app.DeleteFromTransportation) // Изменеие (удаление услуг)
		transportations.PUT("/:transportation_id/user_confirm", app.UserConfirm)                                   // Сформировать создателем
		transportations.PUT("/:transportation_id/moderator_confirm", app.ModeratorConfirm)                         // Завершить отклонить модератором
	}

	r.POST("/api/sign_up", app.Register)
	r.POST("/api/login", app.Login)
	r.POST("/api/logout", app.Logout)
	r.GET("/api/ping", app.WithAuthCheck(role.Moderator), app.Ping)

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

	app.redisClient, err = redis.New(app.config.Redis)
	if err != nil {
		return nil, err
	}

	return &app, nil
}
