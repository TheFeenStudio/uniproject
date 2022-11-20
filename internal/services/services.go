package services

import (
	"uniproject/internal/components"
	"uniproject/internal/services/httpHandler"
)

type Services struct {
	HttpHandler components.HttpServer
}

func NewServices(httpHandler *httpHandler.HttpHandler) *Services {
	return &Services{
		HttpHandler: httpHandler,
	}
}
