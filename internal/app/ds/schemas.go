package ds

import (
	// "gorm.io/gorm"
	"time"
)

type ContainerType struct {
	ContainerTypeId uint   `gorm:"primaryKey"`
	Name            string `gorm:"size:50;not null"`
	Length          int    `gorm:"not null"`
	Height          int    `gorm:"not null"`
	Width           int    `gorm:"not null"`
	MaxGross        int    `gorm:"not null"`
}

type Container struct {
	ContainerID    string `gorm:"primaryKey;size:11;not null"`
	TypeID         uint   `gorm:"not null"`
	ImageURL       string `gorm:"size:100;not null"`
	Decommissioned bool   `gorm:"not null"`

	ContainerType ContainerType `gorm:"ContainerTypeId"`
}

type Status struct {
	StatusID uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:50;not null"`
}

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Login    string `gorm:"size:30;not null"`
	Password string `gorm:"size:30;not null"`
}

type Transportation struct {
	TransportationID uint       `gorm:"primaryKey"`
	StatusID         uint       `gorm:"not null"`
	CreationDate     time.Time  `gorm:"not null;type:date"`
	FormationDate    *time.Time `gorm:"type:date"`
	CompletionDate   *time.Time `gorm:"type:date"`
	Moderator        string     `gorm:"size:50;not null"`
	TransportVehicle string     `gorm:"size:50;not null"`

	Status Status `gorm:"foreignKey:StatusID"`
}

type TransportationComposition struct {
	ContainerID      string `gorm:"primaryKey;size:11;not null"`
	TransportationID uint   `gorm:"primaryKey;not null"`
	Cargo            string `gorm:"size:50;not null"`
	Weight           int    `gorm:"not null"`

	Container      Container      `gorm:"foreignKey:ContainerId"`
	Transportation Transportation `gorm:"foreignKey:TransportationId"`
}
