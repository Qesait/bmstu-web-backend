package schemes

import (
	"bmstu-web-backend/internal/app/ds"
)

type AllContainerTypesResponse struct {
	ContainerTypes []ds.ContainerType `json:"container_types"`
}

type GetContainerResponse struct {
	Container ds.Container `json:"container"`
}

type GetTypeResponse struct {
	ContainerType ds.ContainerType `json:"container_type"`
}

type AllContainersResponse struct {
	Containers []ds.Container `json:"containers"`
}

type AllTransportationsResponse struct {
	Transportations []ds.Transportation `json:"transportations"`
}

type TransportationResponse struct {
	Transportation ds.Transportation `json:"transportation"`
	Containers     []ds.Container    `json:"containers"`
}
