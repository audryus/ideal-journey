package controller

import (
	v1 "ideal-journey/controller/v1"

	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()
	app.Use(apmfiber.Middleware())
	h.api(app)
	return app
}

func (h *Handler) api(app *fiber.App) {
	v1 := v1.NewHandler()
	api := app.Group("/api")
	{
		v1.Init(api)
	}
}
