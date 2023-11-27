package ds

import (
	"time"
)

type User struct {
	UserId    uint   `gorm:"primary_key"`
	Login     string `gorm:"size:30;not null"`
	Password  string `gorm:"size:30;not null"`
	Name      string `gorm:"size:50;not null"`
	Moderator bool   `gorm:"not null"`
}

type Container struct {
	ContainerId uint    `gorm:"primaryKey;not null;"`
	Marking     string  `gorm:"size:11;not null"`
	Type        string  `gorm:"size:50;not null"`
	Length      int     `gorm:"not null"`
	Height      int     `gorm:"not null"`
	Width       int     `gorm:"not null"`
	ImageURL    *string `gorm:"size:100"`
	IsDeleted   bool    `gorm:"not null;default:false"`
	Cargo       string  `gorm:"size:50;not null"`
	Weight      int     `gorm:"not null"`
}

const DRAFT string = "черновик"
const FORMED string = "сформирован"
const COMPELTED string = "завершён"
const REJECTED string = "отклонён"
const DELETED string = "удалён"

type Transportation struct {
	TransportationId uint       `gorm:"primaryKey"`
	Status           string     `gorm:"size:20;not null"`
	CreationDate     time.Time  `gorm:"not null;type:timestamp"`
	FormationDate    *time.Time `gorm:"type:timestamp"`
	CompletionDate   *time.Time `gorm:"type:timestamp"`
	ModeratorId      *string    `json:"-"`
	CustomerId       string     `gorm:"not null"`
	Transport        string     `gorm:"size:50;not null"`

	Moderator *User
	Customer  User
}

type TransportationComposition struct {
	TransportationId string `gorm:"type:uuid"`
	ContainerId      string `gorm:"type:uuid"`

	Container      *Container      `gorm:"foreignKey:ContainerId"`
	Transportation *Transportation `gorm:"foreignKey:TransportationId"`
}
