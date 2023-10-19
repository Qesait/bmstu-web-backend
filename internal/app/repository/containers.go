package repository

import (
	"strings"

	// "gorm.io/gorm"
	"bmstu-web-backend/internal/app/ds"
)

func (r *Repository) GetContainerByID(id string) (*ds.Container, error) {
	container := &ds.Container{UUID: id}
	err := r.db.Preload("ContainerType").
		First(container).
		Error
	if err != nil {
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

func (r *Repository) SaveContainer(container *ds.Container) error {
	err := r.db.Save(container).Error
	// err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&container).Error
	if err != nil {
		return err
	}
	return nil
}
