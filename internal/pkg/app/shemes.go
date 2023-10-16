package app

// import (
// 	""
// )

type Response struct {
}

type AddToTransportationRequest struct {
	ContainerId string `json:"container_id"`
}

type UpdateTransportationRequest struct {
	TransportationId string `json:"transportation_id"`
}