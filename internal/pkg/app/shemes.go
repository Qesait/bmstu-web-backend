package app

import (
	"github.com/google/uuid"
)

type Response struct {
}

type AddToTransportationRequest struct {
	ContainerId uuid.UUID `json:"container_id"`
}

type UpdateTransportationRequest struct {
	TransportationId string `json:"transportation_id"`
}