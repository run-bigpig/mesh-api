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
	manageRouter(app)
}

func v1Router(app *fiber.App) {
	v1 := app.Group("/v1")
	v1.Use(middleware.ErrorIntercept())
	v1.Post("/chat/completions", relay.TextRelayHandler)
	v1.Post("/completions", relay.TextRelayHandler)
	v1.Post("/images/generations", relay.ImageHandler)
}

func manageRouter(app *fiber.App) {
	m := app.Group("/manage")
	m.Use(requestid.New(), middleware.ErrorIntercept())
	m.Get("/adapter/list", manage.ListAdapters)
	m.Post("/model/add", manage.AddModel)
	m.Post("/model/delete", manage.DeleteModel)
	m.Post("/model/update", manage.UpdateModel)
	m.Post("/model/list", manage.ListModel)
	m.Post("/model/detail", manage.FindModel)
	m.Post("/model/setLine", manage.SetModelLine)
	m.Post("/line/add", manage.AddLine)
	m.Post("/line/update", manage.UpdateLine)
	m.Post("/line/detail", manage.FindLine)
	m.Post("/line/list", manage.ListLine)
	m.Post("/line/delete", manage.DeleteLine)
}
