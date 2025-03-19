package handler

import "github.com/NonthapatKim/many_tooth_backend_api/internal/core/port"

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}

type Handler interface{}
