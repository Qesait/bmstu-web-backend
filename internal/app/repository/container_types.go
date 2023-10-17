package repository

import (
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
