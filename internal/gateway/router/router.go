package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler"
)

func InitRouter(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/chat/completions", handler.TextRelayHandler)
	v1.Post("/completions", handler.TextRelayHandler)
	v1.Post("/images/generations", handler.ImageHandler)
}
