package repository

import (
	"bmstu-web-backend/internal/app/ds"
)


func (r *Repository) AddUser(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{
		Login: "login",
	}
	if err := r.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}