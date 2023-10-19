package schemes

import (
	"time"
)

type ContainerRequest struct {
	ContainerId string `uri:"container_id" binding:"uuid"`
}

type GetAllContainersRequest struct {
	ContainerType string `form:"type"`
}

type ChangeContainerRequest struct {
	ContainerId  string     `uri:"container_id" binding:"required,uuid"`
	TypeId       *string    `json:"type_id" binding:"omitempty,uuid"`
	ImageURL     *string    `json:"image_url"`
	PurchaseDate *time.Time `json:"purchase_date"`
	Cargo        *string    `json:"cargo"`
	Weight       *int       `json:"weight"`
	Marking      *string    `json:"marking"`
}

type AddToTransportationRequest struct {
	ContainerId string `json:"container_id" binding:"uuid"`
}

type GetAllTransportationsRequst struct {
	FormationDate *time.Time `form:"formation_date"`
	Status        string     `form:"status"`
}

type TranspostationRequest struct {
	Transpostationid string `uri:"transportation_id" binding:"uuid"`
}

type UpdateTransportationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"uuid"`
	Transport        string `json:"transport"`
}

type DeleteFromTransportationRequest struct {
	Transpostationid string `uri:"transportation_id" binding:"uuid"`
	ContainerId      string `uri:"container_id" binding:"uuid"`
}
