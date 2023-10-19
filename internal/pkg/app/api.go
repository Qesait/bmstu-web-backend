package app

import (
	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/schemes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *Application) getUser() string {
	return "5f58c307-a3f2-4b13-b888-c80ad08d5ed3"
}

func (app *Application) GetAllContainerTypes(c *gin.Context) {
	types, err := app.repo.GetAllContainerTypes()
	if err != nil {
		log.Println("can't get container types from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.AllContainerTypesResponse{ContainerTypes: types})
}

func (app *Application) GetContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		log.Println("can't get container by id", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.GetContainerResponse{Container: *container})
}

func (app *Application) GetAllContainers(c *gin.Context) {
	var request schemes.GetAllContainersRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Println("can't parse request query params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	containers, err := app.repo.GetContainersByType(request.ContainerType)
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) DeleteContainer(c *gin.Context) {
	var request schemes.ContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := app.repo.DeleteContainer(request.ContainerId); err != nil {
		log.Println("can't delete container from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetContainersByType("")
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.AllContainersResponse{Containers: containers})
}

func (app *Application) AddContainer(c *gin.Context) {
	container := &ds.Container{}
	if err := c.BindJSON(container); err != nil {
		log.Println("can't parse requst body", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := app.repo.AddContainer(container); err != nil {
		log.Println("can't save container", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (app *Application) ChangeContainer(c *gin.Context) {
	var request schemes.ChangeContainerRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("can't parse requst body", err)
		c.Status(http.StatusBadRequest)
		return
	}

	container, err := app.repo.GetContainerByID(request.ContainerId)
	if err != nil {
		log.Println("can't get container by id", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if request.TypeId != nil {
		container.TypeId = *request.TypeId
		containerType, err := app.repo.GetContainerType(*request.TypeId)
		if err != nil {
			log.Printf("can't get new container type %v", err)
			c.Status(http.StatusBadRequest)
			return
		}
		container.ContainerType = *containerType
	}
	if request.ImageURL != nil {
		container.ImageURL = *request.ImageURL
	}
	if request.PurchaseDate != nil {
		container.PurchaseDate = *request.PurchaseDate
	}
	if request.Cargo != nil {
		container.Cargo = *request.Cargo
	}
	if request.Weight != nil {
		container.Weight = *request.Weight
	}
	if request.Marking != nil {
		container.Marking = *request.Marking
	}

	log.Println(*container)

	if err := app.repo.SaveContainer(container); err != nil {
		log.Printf("can't save container %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) AddToTranspostation(c *gin.Context) {
	var request schemes.AddToTransportationRequest
	if err := c.BindJSON(&request); err != nil {
		log.Println("can't parse requst body", err)
		c.Status(http.StatusBadRequest)
		return
	}
	var err error

	var transportation *ds.Transportation
	transportation, err = app.repo.GetEditableTransportation(app.getUser())
	if err != nil {
		log.Println("can't get transportation from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if err = app.repo.AddToTransportation(transportation.UUID, request.ContainerId); err != nil {
		log.Println("can't add container to transportation in db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	var containers []ds.Container
	containers, err = app.repo.GetTransportatioinComposition(transportation.UUID)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) GetAllTransportations(c *gin.Context) {
	var request schemes.GetAllTransportationsRequst
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Println("can't parse request query params", err)
		c.Status(http.StatusBadRequest)
		return
	}
	log.Println(request)

	transportations, err := app.repo.GetAllTransportations(request.FormationDate, request.Status)
	if err != nil {
		log.Println("can't get containers from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	log.Println(transportations)
	c.JSON(http.StatusOK, schemes.AllTransportationsResponse{Transportations: transportations})
}

func (app *Application) TranspostationComposition(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	transportation, err := app.repo.GetTransportationById(request.Transpostationid, app.getUser())
	if err != nil {
		log.Printf("can't get transportation by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.Transpostationid)
	if err != nil {
		log.Printf("can't get transportation composition by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) UpdateTransportation(c *gin.Context) {
	var request schemes.UpdateTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("can't parse request body", err)
		c.Status(http.StatusBadRequest)
		return
	}
	transportation, err := app.repo.GetTransportationById(request.TransportationId, app.getUser())
	if err != nil {
		log.Printf("can't get transportation by id %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	transportation.Transport = request.Transport
	if app.repo.SaveTransportation(transportation); err != nil {
		log.Printf("can't save transportation %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	containers, err := app.repo.GetTransportatioinComposition(request.TransportationId)
	if err != nil {
		log.Println("can't get transportation composition from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, schemes.TransportationResponse{Transportation: *transportation, Containers: containers})
}

func (app *Application) DeleteTransportation(c *gin.Context) {
	var request schemes.TranspostationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := app.repo.DeleteTransportation(request.Transpostationid, app.getUser())
	if err != nil {
		log.Println("can't delete transportation from db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (app *Application) DeleteFromTransportation(c *gin.Context) {
	var request schemes.DeleteFromTransportationRequest
	if err := c.ShouldBindUri(&request); err != nil {
		log.Println("can't parse request path params", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if err := app.repo.DeleteFromTransportation(request.Transpostationid, request.ContainerId); err != nil {
		log.Println("can't delete container from transportation in db", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
