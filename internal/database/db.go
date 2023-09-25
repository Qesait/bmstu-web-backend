package db

import (
	"bmstu-web-backend/internal/models"
)

func SetupDB () map[string]models.Container {
	return map[string]models.Container{
		"AAAU1234560": {
			Id: "AAAU1234560",
			Type:         models.DRY_FREIGHT_20,
			ImageURL:     "http://localhost:8080/image/0.jpeg",
			Cargo: models.Cargo{
				Name:   "Бумага",
				Weight: 20000,
			},
			CurrentLocation: "Архангельск",
		},
		"BBBU6543210": {
			Id: "BBBU6543210",
			Type:         models.DRY_FREIGHT_40,
			ImageURL:     "http://localhost:8080/image/1.jpg",
			Cargo: models.Cargo{
				Name:   "Телевизоры",
				Weight: 19000,
			},
			CurrentLocation: "Калининград",
		},
		"CCCU6543210": {
			Id: "CCCU6543210",
			Type:         models.DRY_FREIGHT_20,
			ImageURL:     "http://localhost:8080/image/2.jpg",
			Cargo: models.Cargo{
				Name:   "Зерно",
				Weight: 15000,
			},
			CurrentLocation: "Азов",
		},
		"DDDU6543210": {
			Id: "DDDU6543210",
			Type:         models.HIGH_CUBE_20,
			ImageURL:     "http://localhost:8080/image/4.jpg",
			Cargo: models.Cargo{
				Name:   "Фрукты",
				Weight: 13000,
			},
			CurrentLocation: "Владивосток",
		},
	}
}