package handler

import (
	"time"

	"github.com/NonthapatKim/many-tooh-backend-api/internal/core/domain/response"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) UserLogOut(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HTTPOnly: true,
		Secure:   false,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	result, err := h.svc.UserLogout()
	if err != nil {
		return response.JSONErrorResponse(c, fiber.StatusUnauthorized, err.Error(), nil)
	}

	return response.JSONSuccessResponse(c, result)
}
