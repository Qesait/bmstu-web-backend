package schemes

import (
	"bmstu-web-backend/internal/app/ds"
)

type GetContainerResponse struct {
	Container ds.Container `json:"container"`
}

type AllContainersResponse struct {
	Containers []ds.Container `json:"containers"`
}

type TransportationResponse struct {
	Transportation ds.Transportation `json:"transportation"`
	Containers     []ds.Container    `json:"containers"`
}
