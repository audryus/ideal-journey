package v1

import (
	"ideal-journey/clients/errors"
	"ideal-journey/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) validate(v1 fiber.Router) {
	v1.Post("/auth/validate", func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()
		token := header["Authorization"]

		if strings.TrimSpace(token) == "" {
			return c.Status(401).SendString("")
		}
		payload := &dto.Authenticate{}
		if err := c.BodyParser(&payload); err != nil {
			return errors.InternalServerError("Erro", err)
		}

		if err := h.services.Authenticate.Validate(
			strings.Replace(token, "Bearer ", "", 1),
			payload.Fingerprint); err != nil {
			return c.Status(412).SendString("")
		}

		return c.Status(200).SendString("")
	})
}

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
