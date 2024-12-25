package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler/manage"
	"github.com/run-bigpig/mesh-api/internal/gateway/handler/relay"
	"github.com/run-bigpig/mesh-api/internal/gateway/middleware"
)

func InitRouter(app *fiber.App) {
	app.Use(logger.New())
	v1Router(app)
	adminRouter(app)
}

func v1Router(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Post("/chat/completions", relay.TextRelayHandler)
	v1.Post("/completions", relay.TextRelayHandler)
	v1.Post("/images/generations", relay.ImageHandler)
}

func adminRouter(app *fiber.App) {
	admin := app.Group("/admin")
	admin.Use(requestid.New(), middleware.ErrorIntercept())
	admin.Get("/adapter/list", manage.ListAdapters)
}
