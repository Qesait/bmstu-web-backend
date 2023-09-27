package app

import (
	"bmstu-web-backend/internal/app/config"
	"bmstu-web-backend/internal/app/repository"
)

type Application struct {
	router
	repository
	config *config.Config
}

func (a *Application) Run () {}

// Создание объекта Application, заполнение его конфигом, роутером веб сервера, подключением к базе данных.
func New() (*Application, error) {
	var err error
	app := Application{}
	app.config, err = config.NewConfig()
	if err != nil {
		return nil, err
	}

	return &app, nil
}