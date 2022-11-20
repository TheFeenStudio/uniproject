package application

import (
	"uniproject/internal/config"
	"uniproject/internal/handler"
	"uniproject/internal/services"
	"uniproject/internal/services/httpHandler"
	"uniproject/pkg/logger"
)

type App struct {
	handler *handler.Handler
}

func NewApp() *App {
	return &App{}
}

func (a *App) Setup() {
	cfg := config.LoadConfig()
	newHttpHandler := httpHandler.NewHttpHandler()
	handlerServices := services.NewServices(newHttpHandler)

	a.handler = handler.NewHandler(cfg, handlerServices)
}

func (a *App) Run() {
	a.Setup()
	if err := a.handler.Run(); err != nil {
		logger.Error.Fatalf("[App] error to run handler: %s", err)
	}
}

func (a *App) Stop() {
	if err := a.handler.Stop(); err != nil {
		logger.Error.Fatalf("[App] error to stop handler services: %s", err)
	}
}
