package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"bmstu-web-backend/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}


func (r *Repository) Register(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{
		Login: "login",
	}

	err := r.db.First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}