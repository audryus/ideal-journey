package v1

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(api fiber.Router) {
	v1 := api.Group("/v1")
	h.authenticate(v1)
}
