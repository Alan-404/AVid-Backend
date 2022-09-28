package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	// controllers
	accountController := controllers.NewAccountController()
	userController := controllers.NewUserController()
	videoController := controllers.NewVideoController()
	// group
	userRoutes := app.Group("/user")
	accountRoutes := app.Group("/account")
	videoRoutes := app.Group("/video")

	// routes
	// #user
	userRoutes.Post("/api", userController.UserApi)
	userRoutes.Get("/auth", userController.Auth)
	userRoutes.Get("/avatar", userController.GetAvatar)
	// #account
	accountRoutes.Use("/auth", accountController.Auth)
	// #video
	videoRoutes.Post("/api", videoController.VideoApi)

}
