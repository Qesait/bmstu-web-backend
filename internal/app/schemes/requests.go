package schemes

import (
	"mime/multipart"
	"time"
)

type ContainerRequest struct {
	ContainerId string `uri:"container_id" binding:"required,uuid"`
}

type TypeRequest struct {
	TypeId string `uri:"type_id" binding:"required,uuid"`
}

type GetAllContainersRequest struct {
	ContainerType string `form:"type"`
}

type ChangeContainerRequest struct {
	ContainerId  string                `uri:"container_id" binding:"required,uuid"`
	TypeId       *string               `form:"type_id" json:"type_id" binding:"omitempty,uuid"`
	Image        *multipart.FileHeader `form:"image"`
	PurchaseDate *time.Time            `form:"purchase_date" json:"purchase_date"`
	Cargo        *string               `form:"cargo" json:"cargo"`
	Weight       *int                  `form:"weight" json:"weight"`
	Marking      *string               `form:"marking" json:"marking"`
}

type AddToTransportationRequest struct {
	ContainerId string `uri:"container_id" binding:"required,uuid"`
}

type GetAllTransportationsRequst struct {
	FormationDate *time.Time `form:"formation_date" time_format:"2006-01-02"`
	Status        string     `form:"status"`
}

type TranspostationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
}

type UpdateTransportationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	Transport        string `json:"transport"`
}

type DeleteFromTransportationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	ContainerId      string `uri:"container_id" binding:"required,uuid"`
}

type UserConfirmRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
}

type ModeratorConfirmRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	Status           string `json:"status"`
}
