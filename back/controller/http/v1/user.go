package v1

import (
	"ideal-journey/clients/errors"
	"ideal-journey/clients/logger"
	"ideal-journey/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) get(v1 fiber.Router) {
	v1.Get("/users", func(c *fiber.Ctx) error {
		address := c.Query("publicAddress")

		if strings.TrimSpace(address) == "" {
			err := errors.BadRequestError("address missing", nil)
			return c.Status(err.Status()).JSON(err)
		}

		user, err := h.services.Authenticate.FindById(address)

		if err != nil {
			logger.Error("Erro", err)
			return c.Status(err.Status()).JSON(err)
		}

		return c.JSON(fiber.Map{
			"publicAddress": user.Id,
			"nonce":         user.Nonce,
		})
	})
}

func (h *Handler) create(v1 fiber.Router) {
	v1.Post("/user", func(c *fiber.Ctx) error {
		payload := &dto.Authenticate{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		user, err := h.services.Authenticate.Create(payload.Address)
		if err != nil {
			return c.Status(err.Status()).JSON(err)
		}
		return c.JSON(fiber.Map{
			"publicAddress": user.Id,
			"nonce":         user.Nonce,
		})
	})
}
