package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/run-bigpig/mesh-api/internal/gateway/router"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Mesh API"})
	router.InitRouter(app)
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
