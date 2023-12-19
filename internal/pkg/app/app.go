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

	_ "bmstu-web-backend/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	api := r.Group("/api")
	{
		// Услуги (контейнеры)
		c := api.Group("/containers")
		{
			c.GET("", app.WithAuthCheck(role.NotAuthorized, role.Customer, role.Moderator), app.GetAllContainers)           // Список с поиском
			c.GET("/:id", app.WithAuthCheck(role.NotAuthorized, role.Customer, role.Moderator), app.GetContainer)           // Одна услуга
			c.DELETE("/:id", app.WithAuthCheck(role.Moderator), app.DeleteContainer)                                        // Удаление
			c.PUT("/:id", app.WithAuthCheck(role.Moderator), app.ChangeContainer)                                           // Изменение
			c.POST("", app.WithAuthCheck(role.Moderator), app.AddContainer)                                                 // Добавление
			c.POST("/:id/add_to_transportation", app.WithAuthCheck(role.Customer, role.Moderator), app.AddToTranspostation) // Добавление в заявку
		}
		// Заявки (перевозки)
		t := api.Group("/transportations")
		{
			t.GET("", app.WithAuthCheck(role.Customer, role.Moderator), app.GetAllTransportations)   // Список (отфильтровать по дате формирования и статусу)
			t.GET("/:id", app.WithAuthCheck(role.Customer, role.Moderator), app.GetTranspostation)   // Одна заявка
			t.PUT("", app.WithAuthCheck(role.Customer, role.Moderator), app.UpdateTransportation)    // Изменение (добавление транспорта)
			t.DELETE("", app.WithAuthCheck(role.Customer, role.Moderator), app.DeleteTransportation) //Удаление
			t.DELETE("/delete_container/:id", app.WithAuthCheck(role.Customer, role.Moderator), app.DeleteFromTransportation) // Изменеие (удаление услуг)
			t.PUT("/user_confirm", app.WithAuthCheck(role.Customer, role.Moderator), app.UserConfirm)                         // Сформировать создателем
			t.PUT("/:id/moderator_confirm", app.WithAuthCheck(role.Moderator), app.ModeratorConfirm)                          // Завершить отклонить модератором
			t.PUT("/:id/delivery", app.Delivery)
		}
		// Пользователи (авторизация)
		u := api.Group("/user")
		{
			u.POST("/sign_up", app.Register)
			u.POST("/login", app.Login)
			u.POST("/logout", app.Logout)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
