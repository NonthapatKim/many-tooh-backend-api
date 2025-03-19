package handler

import (
	"github.com/NonthapatKim/many_tooh_backend_api/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	// User
	CreateStaffUser(c *fiber.Ctx) error
}

type handler struct {
	svc port.Service
}

func New(svc port.Service) Handler {
	return &handler{svc: svc}
}
