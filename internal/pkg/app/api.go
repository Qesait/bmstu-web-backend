package app

import (
	"bmstu-web-backend/internal/app/ds"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) GetContainer(c *gin.Context) {
	id := c.Param("id")

	container, err := app.repo.GetContainerByID(id)
	if err != nil {
		log.Printf("can't get product by id %v", err)
		c.Error(err)
		return
	}

	if len(container.Cargo) == 0 {
		log.Println("empty")
	}

	c.HTML(http.StatusOK, "item-info.tmpl", *container)
}

type GetContainersResponse struct {
	Containers []ds.Container
	Search     string
}

func (app *Application) GetContainers(c *gin.Context) {
	containerType := c.Query("type")

	containers, err := app.repo.GetContainersByType(containerType)
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Error(err)
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", GetContainersResponse{
		Search:     containerType,
		Containers: containers,
	})
}

func (app *Application) DecommissionContainer(c *gin.Context) {
	id := c.PostForm("delete")

	app.repo.DecommissionContainer(id)

	containers, err := app.repo.GetContainersByType("")
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Error(err)
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", GetContainersResponse{
		Search:     "",
		Containers: containers,
	})
}
