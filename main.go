package main

import (
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	configs.ConnectDB()

	routes.Route(app)
	app.Listen(":5200")
}
