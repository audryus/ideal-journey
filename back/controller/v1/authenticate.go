package v1

import "github.com/gofiber/fiber/v2"

func (h *Handler) authenticate(v1 fiber.Router) {
	v1.Get("/authenticate", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
