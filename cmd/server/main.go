package main

import (
	"log"
	"bmstu-web-backend/internal/pkg/app"
)


// TODO: change
// @title Container loginstics
// @version 1.0

// @host 127.0.0.1:8080
// @schemes http
// @BasePath /

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
