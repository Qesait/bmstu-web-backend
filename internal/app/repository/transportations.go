package repository

import (
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"

	"bmstu-web-backend/internal/app/ds"
)

// TODO: Как-то фильтровать по дате формирования
func (r *Repository) GetAllTransportations(_ *time.Time, status string) ([]ds.Transportation, error) {
	var transportations []ds.Transportation

	err := r.db.Where("LOWER(status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Find(&transportations).Error
	if err != nil {
		return nil, err
	}
	return transportations, nil
}

func (r *Repository) DeleteContainer(id string) error {
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

func (r *Repository) GetEditableTransportation(customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.First(transportation, ds.Transportation{Status: "введён", CustomerId: &customerId}).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		transportation = &ds.Transportation{CreationDate: time.Now(), CustomerId: &customerId}
		err := r.db.Create(transportation).Error
		if err != nil {
			return nil, err
		}
	}
	return transportation, nil
}

func (r *Repository) GetTransportationById(transportationId, customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.First(transportation, ds.Transportation{UUID: transportationId, CustomerId: &customerId}).Error
	if err != nil {
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) AddToTransportation(transportationId, containerId string) error {
	tComposition := ds.TransportationComposition{TransportationId: transportationId, ContainerId: containerId}
	err := r.db.Create(&tComposition).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetTransportatioinComposition(transportationId string) ([]ds.Container, error) {

	var containerIDs []string
	err := r.db.Table("transportation_compositions").
		Where("transportation_id = ?", transportationId).
		Pluck("container_id", &containerIDs).Error
	if err != nil {
		return nil, err
	}
	log.Println(containerIDs)
	var containers []ds.Container
	err = r.db.Preload("ContainerType").
		Find(&containers, containerIDs).Error
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *Repository) AddTransport(transportationId string, transport string) error {
	transportation := &ds.Transportation{UUID: transportationId}
	var err error

	err = r.db.First(transportation).Error
	if err != nil {
		return err
	}

	transportation.Transport = transport
	err = r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SaveTransportation(transportation *ds.Transportation) error {
	err := r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteTransportation(transportationId, customerId string) error {
	transportation := &ds.Transportation{UUID: transportationId, CustomerId: &customerId}
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

func (r *Repository) DeleteFromTransportation(transportationId, ContainerId string) error {
	tComposition := &ds.TransportationComposition{TransportationId: transportationId, ContainerId: ContainerId}

	err := r.db.Delete(tComposition).Error
	if err != nil {
		return err
	}
	return nil
}
