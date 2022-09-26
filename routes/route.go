package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// controllers
	accountController := controllers.NewAccountController()
	userController := controllers.NewUserController()
	// group
	userRoutes := app.Group("/user")
	accountRoutes := app.Group("/account")

	// routes
	userRoutes.Post("/api", userController.UserApi)
	userRoutes.Get("/auth", userController.Auth)
	accountRoutes.Post("/auth", accountController.Auth)

}
