package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/run-bigpig/mesh-api/internal/config"
	"github.com/run-bigpig/mesh-api/internal/data/driver"
	"github.com/run-bigpig/mesh-api/internal/gateway/router"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()
	config.Set(*configFile)
	app := fiber.New(fiber.Config{AppName: config.Get().App})
	router.InitRouter(app)
	driver.NewMySQL()
	log.Fatal(app.Listen(config.Get().Listen))
}
