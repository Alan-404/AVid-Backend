package controllers

import (
	"context"
	"net/http"
	"server/dto"
	"server/models"
	"server/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	userService    services.UserService
	accountService services.AccountService
}

func NewUserController() *UserController {
	userController := new(UserController)
	userController.userService = *services.NewUserService()
	userController.accountService = *services.NewAccountService()
	return userController
}

func (userController *UserController) UserApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	var userData *dto.CreateUserDTO

	if err := c.BodyParser(&userData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	user := models.User{
		Id:        primitive.NewObjectID(),
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Phone:     userData.Phone,
		BDate:     userData.BDate,
		Gender:    userData.Gender,
	}

	newUserId := userController.userService.CreateUser(ctx, user)

	if newUserId == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	account := models.Account{
		UserId:   *newUserId,
		Password: userData.Password,
		Role:     userData.Role,
	}

	addedAccount := userController.accountService.CreateAccount(ctx, account)

	if addedAccount == nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateUserDTO{Success: true, User: user})
}
