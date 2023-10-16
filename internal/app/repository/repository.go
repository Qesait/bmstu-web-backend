package repository

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
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

func (r *Repository) GetContainerByID(id uuid.UUID) (*ds.Container, error) {
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

func (r *Repository) DeleteContainer(id uuid.UUID) error {
	container := &ds.Container{UUID: id}
	var err error

	err = r.db.First(container).Error
	if err != nil {
		return err
	}
	container.IsDeleted = true
	err = r.db.Save(container).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetEditableTransportation() (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.First(transportation, ds.Transportation{Status: "введён"}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		transportation = &ds.Transportation{CreationDate: time.Now()}
		err := r.db.Create(transportation).Error
		if err != nil {
			return nil, err
		}
	}
	return transportation, nil
}

func (r *Repository) AddToTransportation(transportationId, containerId uuid.UUID) error {
	tComposition := ds.TransportationComposition{TransportationId: transportationId, ContainerId: containerId}
	err := r.db.Create(&tComposition).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTransportatioinComposition(transportationId uuid.UUID) ([]ds.TransportationComposition, error) {
	var containers []ds.TransportationComposition

	err := r.db.Preload("Container.ContainerType").
		Find(&containers, ds.TransportationComposition{TransportationId: transportationId}).Error
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *Repository) AddTransportVehicle(transportationId uuid.UUID, transport string) error {
	transportation := &ds.Transportation{UUID: transportationId}
	var err error

	err = r.db.First(transportation).Error
	if err != nil {
		return err
	}

	transportation.TransportVehicle = transport
	err = r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTransportation(transportationId uuid.UUID) error {
	transportation := &ds.Transportation{UUID: transportationId}
	var err error

	err = r.db.First(transportation).Error
	if err != nil {
		return err
	}

	transportation.Status = "удалён"
	err = r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromTransportation(transportationId, ContainerId uuid.UUID) error {
	tComposition := &ds.TransportationComposition{TransportationId: transportationId, ContainerId: ContainerId}
	var err error

	err = r.db.First(tComposition).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(tComposition).Error
	if err != nil {
		return err
	}
	return nil
}
