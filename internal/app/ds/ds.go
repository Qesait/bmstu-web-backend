package ds

import (
	"time"
)

type User struct {
	UUID      string `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Login     string `gorm:"size:30;not null"`
	Password  string `gorm:"size:30;not null"`
	Name      string `gorm:"size:50;not null"`
	Moderator bool   `gorm:"not null"`
}

type Container struct {
	UUID      string  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid" binding:"-"`
	Marking   string  `gorm:"size:11;not null" form:"marking" json:"marking" binding:"required,max=11"`
	Type      string  `gorm:"size:50;not null" form:"type" json:"type" binding:"required,max=50"`
	Length    int     `gorm:"not null" form:"length" json:"length" binding:"required"`
	Height    int     `gorm:"not null" form:"height" json:"height" binding:"required"`
	Width     int     `gorm:"not null" form:"width" json:"width" binding:"required"`
	ImageURL  *string `gorm:"size:100" json:"image_url" binding:"-"`
	IsDeleted bool    `gorm:"not null;default:false" json:"-" binding:"-"`
	Cargo     string  `gorm:"size:50;not null" form:"cargo" json:"cargo" binding:"required,max=50"`
	Weight    int     `gorm:"not null" form:"weight" json:"weight" binding:"required"`
}

const DRAFT string = "черновик"
const FORMED string = "сформирован"
const COMPELTED string = "завершён"
const REJECTED string = "отклонён"
const DELETED string = "удалён"

type Transportation struct {
	UUID           string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"uuid"`
	Status         string     `gorm:"size:20;not null" json:"status"`
	CreationDate   time.Time  `gorm:"not null;type:date" json:"creation_date"`
	FormationDate  *time.Time `gorm:"type:date" json:"formation_date"`
	CompletionDate *time.Time `gorm:"type:date" json:"completion_date"`
	ModeratorId    *string    `json:"moderator_id"`
	CustomerId     string     `gorm:"not null" json:"customer_id"`
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
