package ds

import (
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

type Status struct {
	StatusId uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:50;not null"`
}

type User struct {
	UserId    uint   `gorm:"primaryKey"`
	Login     string `gorm:"size:30;not null"`
	Password  string `gorm:"size:30;not null"`
	Name      string `gorm:"size:50;not null"`
	Moderator bool   `gorm:"not null"`
}

type Container struct {
	ContainerId    string    `gorm:"primaryKey;size:11;not null;autoIncrement:false"`
	TypeId         uint      `gorm:"not null"`
	ImageURL       string    `gorm:"size:100;not null"`
	Decommissioned bool      `gorm:"not null"`
	PurchaseDate   time.Time `gorm:"not null;type:date"`

	ContainerType ContainerType `gorm:"preload:false;foreignKey:TypeId"`
}

type Transportation struct {
	TransportationId uint       `gorm:"primaryKey"`
	StatusId         uint       `gorm:"not null"`
	CreationDate     time.Time  `gorm:"not null;type:date"`
	FormationDate    *time.Time `gorm:"type:date"`
	CompletionDate   *time.Time `gorm:"type:date"`
	ModeratorId      uint       `gorm:"not null"`
	CustomerId       uint       `gorm:"not null"`
	TransportVehicle string     `gorm:"size:50;not null"`

	Status    Status
	Moderator User `gorm:"foreignKey:ModeratorId"`
	Customer  User `gorm:"foreignKey:CustomerId"`
}

type TransportationComposition struct {
	ContainerId      string `gorm:"primaryKey;size:11;not null;autoIncrement:false"`
	TransportationId uint   `gorm:"primaryKey;not null;autoIncrement:false"`
	Cargo            string `gorm:"size:50;not null"`
	Weight           int    `gorm:"not null"`

	Container      *Container      `gorm:"foreignKey:ContainerId"`
	Transportation *Transportation `gorm:"foreignKey:TransportationId"`
}
