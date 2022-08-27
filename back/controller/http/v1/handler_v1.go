package v1

import (
	"ideal-journey/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *usecase.Services
}

func NewHandler(services *usecase.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api fiber.Router) {
	v1 := api.Group("/v1")

	h.authenticate(v1)
	h.get(v1)
	h.create(v1)

}
