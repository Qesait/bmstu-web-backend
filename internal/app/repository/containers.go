package repository

import (
	"strings"
	"errors"

	"gorm.io/gorm"
	"bmstu-web-backend/internal/app/ds"
)

func (r *Repository) GetContainerByID(id string) (*ds.Container, error) {
	container := &ds.Container{UUID: id}
	err := r.db.First(container, "is_deleted = ?", false).Error
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

	err := r.db.Where("LOWER(type) LIKE ?", "%"+strings.ToLower(containerType)+"%").
		Where("is_deleted = ?", false).
		Find(&containers).Error
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *Repository) SaveContainer(container *ds.Container) error {
	err := r.db.Save(container).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) AddToTransportation(transportationId, containerId string) error {
	tComposition := ds.TransportationComposition{TransportationId: transportationId, ContainerId: containerId}
	err := r.db.Create(&tComposition).Error
	if err != nil {
		return err
	}
	return nil
}