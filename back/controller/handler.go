package controller

import (
	v1 "ideal-journey/controller/http/v1"
	"ideal-journey/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.elastic.co/apm/module/apmfiber/v2"
)

type Handler struct {
	services *usecase.Services
}

func NewHandler(services *usecase.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *fiber.App {
	app := fiber.New()
	app.Use(apmfiber.Middleware())
	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	h.api(app)
	return app
}

func (h *Handler) api(app *fiber.App) {
	v1 := v1.NewHandler(h.services)
	api := app.Group("/api")
	{
		v1.Init(api)
	}
}
