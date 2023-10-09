package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bmstu-web-backend/internal/app/ds"
	"bmstu-web-backend/internal/app/dsn"
)

func main() {
	_ = godotenv.Load()
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&ds.User{},
		&ds.ContainerType{},
		&ds.Container{},
		&ds.Status{},
		&ds.Transportation{},
		&ds.TransportationComposition{},
	)
	if err != nil {
		panic("cant migrate db")
	}
}