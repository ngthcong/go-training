package app

import (
	"school-mongo/internal/handler"
)

type (
	App struct {
		handler handler.Handler
	}
)

func New(handler handler.Handler) *App {
	return &App{handler: handler}
}

func (a App) Start() {
	a.handler.InitProgram()
}
