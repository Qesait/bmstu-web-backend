package ds

import (
	"time"

	"github.com/google/uuid"
)

type ContainerType struct {
	UUID     uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name     string    `gorm:"size:50;not null"`
	Length   int       `gorm:"not null"`
	Height   int       `gorm:"not null"`
	Width    int       `gorm:"not null"`
	MaxGross int       `gorm:"not null"`
}

type User struct {
	UUID      uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Login     string    `gorm:"size:30;not null"`
	Password  string    `gorm:"size:30;not null"`
	Name      string    `gorm:"size:50;not null"`
	Moderator bool      `gorm:"not null"`
}

type Container struct {
	UUID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	TypeId       uuid.UUID `gorm:"type:uuid;not null"`
	ImageURL     string    `gorm:"size:100;not null"`
	IsDeleted    bool      `gorm:"not null"`
	PurchaseDate time.Time `gorm:"not null;type:date"`
	Cargo        string    `gorm:"size:50;not null"`
	Weight       int       `gorm:"not null"`
	Marking      string    `gorm:"size:11;not null"`

	ContainerType ContainerType `gorm:"preload:false;foreignKey:TypeId"`
}

type Transportation struct {
	UUID             uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Status           string     `gorm:"size:20;not null;default:'введён'"` // введён, в работе, завершён, отменён, удалён
	CreationDate     time.Time  `gorm:"not null;type:date"`
	FormationDate    *time.Time `gorm:"type:date"`
	CompletionDate   *time.Time `gorm:"type:date"`
	ModeratorId      *uuid.UUID
	CustomerId       *uuid.UUID
	TransportVehicle string `gorm:"size:50;not null"`

	Moderator User `gorm:"foreignKey:ModeratorId"`
	Customer  User `gorm:"foreignKey:CustomerId"`
}

type TransportationComposition struct {
	TransportationId uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ContainerId      uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	Container      *Container      `gorm:"foreignKey:ContainerId"`
	Transportation *Transportation `gorm:"foreignKey:TransportationId"`
}
