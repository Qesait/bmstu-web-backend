package schemes

import (
	"bmstu-web-backend/internal/app/ds"
	"mime/multipart"
	"time"
)

type ContainerRequest struct {
	ContainerId string `uri:"container_id" binding:"required,uuid"`
}

type GetAllContainersRequest struct {
	ContainerType string `form:"type"`
}

type AddContainerRequest struct {
	ds.Container
	Image *multipart.FileHeader `form:"image" json:"image"`
}

type ChangeContainerRequest struct {
	ContainerId string                `uri:"container_id" binding:"required,uuid"`
	Marking     *string               `form:"marking" json:"marking" binding:"omitempty,max=11"`
	Type        *string               `form:"type" json:"type" binding:"omitempty,max=50"`
	Length      *int                  `form:"length" json:"length"`
	Height      *int                  `form:"height" json:"height"`
	Width       *int                  `form:"width" json:"width"`
	Image       *multipart.FileHeader `form:"image" json:"image"`
	Cargo       *string               `form:"cargo" json:"cargo" binding:"omitempty,max=50"`
	Weight      *int                  `form:"weight" json:"weight"`
}

type AddToTransportationRequest struct {
	ContainerId string `uri:"container_id" binding:"required,uuid"`
}

type GetAllTransportationsRequst struct {
	FormationDateStart *time.Time `form:"formation_date_start" json:"formation_date_start" time_format:"2006-01-02 15:04:05"`
	FormationDateEnd   *time.Time `form:"formation_date_end" json:"formation_date_end" time_format:"2006-01-02 15:04:05"`
	Status             string     `form:"status" json:"status"`
}

type TranspostationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
}

type UpdateTransportationRequest struct {
	URI struct {
		TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	}
	Transport string `form:"transport" json:"transport" binding:"required,max=50"`
}

type DeleteFromTransportationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	ContainerId      string `uri:"container_id" binding:"required,uuid"`
}

type UserConfirmRequest struct {
	Confirm bool `form:"confirm" binding:"required"`
}

type ModeratorConfirmRequest struct {
	URI struct {
		TransportationId string `uri:"transportation_id" binding:"required,uuid"`
	}
	Confirm bool `form:"confirm" binding:"required"`
}

type LoginReq struct {
	Login    string `form:"login" binding:"required,max=30"`
	Password string `form:"password" binding:"required,max=30"`
}

type RegisterReq struct {
	Login    string `form:"login" binding:"required,max=30"`
	Password string `form:"password" binding:"required,max=30"`
}