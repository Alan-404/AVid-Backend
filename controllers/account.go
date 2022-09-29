package controllers

import (
	"context"
	"net/http"
	"server/dto"
	"server/middleware"
	"server/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountController struct {
	accountService *services.AccountService
	userService    *services.UserService
}

func NewAccountController() *AccountController {
	accountController := new(AccountController)
	accountController.accountService = services.NewAccountService()
	accountController.userService = services.NewUserService()
	return accountController
}

func (accountController *AccountController) Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	var loginData *dto.LoginDTO

	err := c.BodyParser(&loginData)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	user := accountController.userService.GetUserByEmail(ctx, loginData.Email)

	if user == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	account := accountController.accountService.GetAccountByUserId(ctx, user.Id)

	if account == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	checkPassword := accountController.accountService.CheckPassword(account, loginData.Password)

	if checkPassword == false {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	accessToken := middleware.GenerateToken(account.Id)

	if accessToken == "" {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	return c.Status(http.StatusAccepted).JSON(dto.ResponseLoginDTO{Success: true, AccessToken: accessToken})
}

func (accountController *AccountController) ChangePassword(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	var changePasswordData *dto.ChangePasswordDTO

	err := c.BodyParser(&changePasswordData)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}
	accountId, _ := primitive.ObjectIDFromHex(c.GetReqHeaders()["Id"])

	account := accountController.accountService.GetAccountById(ctx, accountId)

	if account == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	changePassword := accountController.accountService.ChangePassword(ctx, account, changePasswordData.OldPassword, changePasswordData.NewPassword)

	if changePassword == false {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseLoginDTO{Success: false})
	}

	return c.Status(http.StatusAccepted).JSON(dto.ResponseLoginDTO{Success: true})
}
