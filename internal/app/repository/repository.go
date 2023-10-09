package repository

import (
	"strings"

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

func (r *Repository) GetContainerByID(id string) (*ds.Container, error) {
	container := &ds.Container{}
	err := r.db.Where("container_id = ?", id).
		Preload("ContainerType").First(container).Error
	if err != nil {
		return nil, err
	}

	return container, nil
}

func (r *Repository) GetContainersByType(containerType string) ([]ds.Container, error) {
	var containers []ds.Container

	err := r.db.Preload("ContainerType").
		Joins("JOIN container_types ON containers.type_id = container_types.container_type_id").
		Where("LOWER(container_types.name) LIKE ?", "%"+strings.ToLower(containerType)+"%").
		Where("is_deleted = ?", false).
		Find(&containers).Error

	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (r *Repository) DecommissionContainer(id string) error {
	err := r.db.Exec("UPDATE containers SET is_deleted = ? WHERE container_id = ?", true, id).Error
	if err != nil {
		return err
	}

	return nil
}
