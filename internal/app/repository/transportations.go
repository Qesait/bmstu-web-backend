package repository

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"bmstu-web-backend/internal/app/ds"
)

func (r *Repository) GetAllTransportations(customerId *string, formationDateStart, formationDateEnd *time.Time, status string) ([]ds.Transportation, error) {
	var transportations []ds.Transportation

	query := r.db.Preload("Customer").Preload("Moderator").
		Where("LOWER(status) LIKE ?", "%"+strings.ToLower(status)+"%").
		Where("status != ? AND status != ?", ds.DELETED, ds.DRAFT)

	if customerId != nil {
		query = query.Where("customer_id = ?", *customerId)
	}
	if formationDateStart != nil && formationDateEnd != nil {
		query = query.Where("formation_date BETWEEN ? AND ?", *formationDateStart, *formationDateEnd)
	} else if formationDateStart != nil {
		query = query.Where("formation_date >= ?", *formationDateStart)
	} else if formationDateEnd != nil {
		query = query.Where("formation_date <= ?", *formationDateEnd)
	}

	if err := query.Find(&transportations).Error; err != nil {
		return nil, err
	}
	return transportations, nil
}

func (r *Repository) GetDraftTransportation(customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	err := r.db.First(transportation, ds.Transportation{Status: ds.DRAFT, CustomerId: customerId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) CreateDraftTransportation(customerId string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{CreationDate: time.Now(), CustomerId: customerId, Status: ds.DRAFT}
	err := r.db.Create(transportation).Error
	if err != nil {
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) GetTransportationById(transportationId string, userId *string) (*ds.Transportation, error) {
	transportation := &ds.Transportation{}
	query := r.db.Preload("Moderator").Preload("Customer").
		Where("status != ?", ds.DELETED)
	if userId != nil {
		query = query.Where("customer_id = ?", userId)
	}
	err := query.First(transportation, ds.Transportation{UUID: transportationId}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return transportation, nil
}

func (r *Repository) GetTransportatioinComposition(transportationId string) ([]ds.Container, error) {
	var containers []ds.Container

	err := r.db.Table("transportation_compositions").
		Select("containers.*").
		Joins("JOIN containers ON transportation_compositions.container_id = containers.uuid").
		Where(ds.TransportationComposition{TransportationId: transportationId}).
		Scan(&containers).Error

	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *Repository) SaveTransportation(transportation *ds.Transportation) error {
	err := r.db.Save(transportation).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteFromTransportation(transportationId, ContainerId string) error {
	err := r.db.Delete(&ds.TransportationComposition{TransportationId: transportationId, ContainerId: ContainerId}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CountContainers(transportationId string) (int64, error) {
	var count int64
	err := r.db.Model(&ds.TransportationComposition{}).
		Where("transportation_id = ?", transportationId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
