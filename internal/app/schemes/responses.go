package schemes

import (
	"bmstu-web-backend/internal/app/ds"
)

type AllContainersResponse struct {
	Containers []ds.Container `json:"containers"`
}

type GetAllContainersResponse struct {
	DraftTransportation *ds.Transportation `json:"draft_transportation"`
	Containers          []ds.Container     `json:"containers"`
}

type AllTransportationsResponse struct {
	Transportations []ds.Transportation `json:"transportations"`
}

type TransportationResponse struct {
	Transportation ds.Transportation `json:"transportation"`
	Containers     []ds.Container    `json:"containers"`
}

type UpdateTransportationResponse struct {
	Transportation ds.Transportation `json:"transportation"`
}
