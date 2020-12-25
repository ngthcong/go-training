package app

import "school-mongo/internal/services"

type (
	App struct {
		service services.Service
	}
)

func New(service services.Service) *App {
	return &App{service: service}
}
