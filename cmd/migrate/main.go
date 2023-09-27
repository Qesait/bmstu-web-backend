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

	// Migrate the schema
	err = db.AutoMigrate(
		&ds.Transportation{},
		&ds.Status{},
		&ds.ContainerType{},
		&ds.User{},
		&ds.Container{},
		&ds.TransportationComposition{},
	)
	if err != nil {
		panic("cant migrate db")
	}
}