package routes

import (
	"server/controllers"
	"server/middleware"

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
	userRoutes.Get("/auth", middleware.GetAccountId, userController.Auth)
	userRoutes.Get("/avatar", userController.GetAvatar)
	// #account
	accountRoutes.Post("/auth", accountController.Login)
	accountRoutes.Put("/auth", middleware.GetAccountId, accountController.ChangePassword)
	// #video
	videoRoutes.Post("/api", middleware.GetAccountId, videoController.VideoApi)
	videoRoutes.Get("/api", videoController.GetVideos)
	videoRoutes.Get("/media", videoController.GetMedia)
}
