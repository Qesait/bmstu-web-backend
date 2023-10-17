package repository

import (
	"errors"
	"gorm.io/gorm"
	"strings"

	"bmstu-web-backend/internal/app/ds"
)

func (r *Repository) GetContainerByID(id string) (*ds.Container, error) {
	container := &ds.Container{UUID: id}
	err := r.db.Preload("ContainerType").
		First(container).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return container, nil
}

func (r *Repository) AddContainer(container *ds.Container) error {
	err := r.db.Create(&container).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetContainersByType(containerType string) ([]ds.Container, error) {
	var containers []ds.Container

	err := r.db.Joins("ContainerType").
		Where("LOWER(name) LIKE ?", "%"+strings.ToLower(containerType)+"%").
		Where("is_deleted = ?", false).
		Find(&containers).Error
	if err != nil {
		return nil, err
	}
	return containers, nil
}