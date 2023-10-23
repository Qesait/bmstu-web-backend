package repository

import (
	"errors"
	"gorm.io/gorm"
	"bmstu-web-backend/internal/app/ds"
)

func (r *Repository) GetAllContainerTypes() ([]ds.ContainerType, error) {
	var types []ds.ContainerType

	err := r.db.Find(&types).Error
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (r *Repository) GetContainerType(id string) (*ds.ContainerType, error) {
	containerType := &ds.ContainerType{UUID: id}

	err := r.db.First(&containerType).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return containerType, nil
}
