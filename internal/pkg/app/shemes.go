package app

import (
	// "github.com/google/uuid"
)

type ContainerRequest struct {
	ContainerId string `uri:"container_id" binding:"uuid"`
}

type GetAllContainersRequest struct {
	ContainerType string `form:"type"`
}

type AddToTransportationRequest struct {
	ContainerId string `json:"container_id"`
}

type TranspostationRequest struct {
	Transpostationid string `uri:"transportation_id" binding:"uuid"`
}

type UpdateTransportationRequest struct {
	TransportationId string `uri:"transportation_id" binding:"uuid"`
	Vehicle          string `json:"transport_vehicle"`
}

type DeleteFromTransportationRequest struct {
	Transpostationid string `uri:"transportation_id" binding:"uuid"`
	ContainerId string `uri:"container_id" binding:"uuid"`
}