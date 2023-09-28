package main

import (
	"log"
	"bmstu-web-backend/internal/pkg/app"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
