package handler

import "github.com/basterrus/x2gobroker/internal/service"

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: services}
}
