package schemes

import (
	"bmstu-web-backend/internal/app/ds"
	"time"
)

type AllContainersResponse struct {
	Containers []ds.Container `json:"containers"`
}

type TransportationShort struct {
	UUID           string `json:"uuid"`
	ContainerCount int    `json:"container_count"`
}

type GetAllContainersResponse struct {
	DraftTransportation *TransportationShort `json:"draft_transportation"`
	Containers          []ds.Container       `json:"containers"`
}

type AllTransportationsResponse struct {
	Transportations []TransportationOutput `json:"transportations"`
}

type TransportationResponse struct {
	Transportation TransportationOutput `json:"transportation"`
	Containers     []ds.Container       `json:"containers"`
}

type UpdateTransportationResponse struct {
	Transportation TransportationOutput `json:"transportation"`
}

type TransportationOutput struct {
	UUID           string  `json:"uuid"`
	Status         string  `json:"status"`
	CreationDate   string  `json:"creation_date"`
	FormationDate  *string `json:"formation_date"`
	CompletionDate *string `json:"completion_date"`
	Moderator      *string `json:"moderator"`
	Customer       string  `json:"customer"`
	Transport      string  `json:"transport"`
}

func ConvertTransportation(transportation *ds.Transportation) TransportationOutput {
	output := TransportationOutput{
		UUID:         transportation.UUID,
		Status:       transportation.Status,
		CreationDate: transportation.CreationDate.Format("2006-01-02 15:04:05"),
		Transport:    transportation.Transport,
		Customer:     transportation.Customer.Login,
	}

	if transportation.FormationDate != nil {
		formationDate := transportation.FormationDate.Format("2006-01-02 15:04:05")
		output.FormationDate = &formationDate
	}

	if transportation.CompletionDate != nil {
		completionDate := transportation.CompletionDate.Format("2006-01-02 15:04:05")
		output.CompletionDate = &completionDate
	}

	if transportation.Moderator != nil {
		output.Moderator = &transportation.Moderator.Login
	}

	return output
}

type LoginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
}

type SwaggerLoginResp struct {
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type RegisterResp struct {
	Ok bool `json:"ok"`
}
