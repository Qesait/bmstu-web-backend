package ds

import (
	"time"
)

type ContainerType struct {
	UUID     string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"-"`
	Name     string `gorm:"size:50;not null" json:"name"`
	Length   int    `gorm:"not null" json:"length"`
	Height   int    `gorm:"not null" json:"height"`
	Width    int    `gorm:"not null" json:"width"`
	MaxGross int    `gorm:"not null" json:"max_gross"`
}

type User struct {
	UUID      string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	Login     string `gorm:"size:30;not null" json:"login"`
	Password  string `gorm:"size:30;not null" json:"password"`
	Name      string `gorm:"size:50;not null" json:"name"`
	Moderator bool   `gorm:"not null" json:"moderator"`
}

type Container struct {
	UUID         string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	TypeId       string    `gorm:"type:uuid;not null" json:"-"`
	ImageURL     string    `gorm:"size:100;not null" json:"image_url"`
	IsDeleted    bool      `gorm:"not null" json:"is_deleted"`
	PurchaseDate time.Time `gorm:"not null;type:date" json:"purchase_date"`
	Cargo        string    `gorm:"size:50;not null" json:"cargo"`
	Weight       int       `gorm:"not null" json:"weight"`
	Marking      string    `gorm:"size:11;not null" json:"marking"`

	ContainerType ContainerType `gorm:"preload:false;foreignKey:TypeId" json:"container_type"`
}

type Transportation struct {
	UUID           string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	Status         string     `gorm:"size:20;not null;default:'введён'" json:"status"` // введён, в работе, завершён, отменён, удалён
	CreationDate   time.Time  `gorm:"not null;type:date" json:"creation_date"`
	FormationDate  *time.Time `gorm:"type:date" json:"formation_date"`
	CompletionDate *time.Time `gorm:"type:date" json:"completion_date"`
	ModeratorId    *string    `json:"moderator_id"`
	CustomerId     *string    `json:"customer_id"`
	Transport      string     `gorm:"size:50;not null" json:"transport"`

	Moderator User `gorm:"foreignKey:ModeratorId" json:"-"`
	Customer  User `gorm:"foreignKey:CustomerId" json:"-"`
}

type TransportationComposition struct {
	TransportationId string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"transportation_id"`
	ContainerId      string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"container_id"`

	Container      *Container      `gorm:"foreignKey:ContainerId" json:"container"`
	Transportation *Transportation `gorm:"foreignKey:TransportationId" json:"transportation"`
}
