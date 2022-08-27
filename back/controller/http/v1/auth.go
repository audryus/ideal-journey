package v1

import (
	"ideal-journey/clients/errors"
	"ideal-journey/dto"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) authenticate(v1 fiber.Router) {
	v1.Post("/auth", func(c *fiber.Ctx) error {
		payload := &dto.Authenticate{}

		if err := c.BodyParser(&payload); err != nil {
			return errors.InternalServerError("Erro", err)
		}
		token, errRes := h.services.Authenticate.Auth(payload)

		if errRes != nil {
			return c.Status(errRes.Status()).JSON(errRes)
		}

		return c.JSON(fiber.Map{
			"access_token": token,
		})
	})
}
