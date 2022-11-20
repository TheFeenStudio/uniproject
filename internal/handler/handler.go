package handler

import (
	"uniproject/internal/config"
	"uniproject/internal/services"
)

type Handler struct {
	cfg      *config.Config
	services *services.Services
}

func NewHandler(cfg *config.Config, services *services.Services) *Handler {
	return &Handler{cfg: cfg, services: services}
}

func (h *Handler) Stop() error {
	//todo: stop handler services here
	return nil
}

func (h *Handler) Run() error {
	if err := h.services.HttpHandler.RunService(h.cfg.ServerPort); err != nil {
		return err
	}

	return nil
}
