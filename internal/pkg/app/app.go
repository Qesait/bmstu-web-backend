package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/dsn"
	"bmstu-web-backend/internal/app/repository"
)

type Application struct {
	repo   *repository.Repository
	config *config.Config
	// dsn string
}

func (a *Application) Run() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/containers/:id", func(c *gin.Context) {
		id := c.Param("id")

		log.Printf("id recived %s\n", id)
		// получаем данные по товару
		container, err := a.repo.GetContainerByID(id)
		if err != nil { // если не получилось
			log.Printf("cant get product by id %v", err)
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"conatiner": container,
		})
		return
	})

	r.GET("/containers", func(c *gin.Context) {
		containerType := c.Query("type") // получаем из запроса query string

		if containerType == "" {
			// получаем данные по товару
			containers, err := a.repo.GetAllContainers()
			if err != nil { // если не получилось
				log.Println("cant get containers", err)
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"containers": containers,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "not implemented",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.Static("/image", "./static/image")
	r.Static("/css", "./static/css")

	r.Run()

	log.Println("Server down")
}

// Создание объекта Application, заполнение его конфигом, роутером веб сервера, подключением к базе данных.
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
