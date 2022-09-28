package main

import (
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{DisablePreParseMultipartForm: true, StreamRequestBody: true})

	app.Use(cors.New())

	configs.ConnectDB()

	routes.Route(app)
	app.Listen(":5200")
}
