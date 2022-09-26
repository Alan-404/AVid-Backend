package controllers

import (
	"context"
	"net/http"
	"server/dto"
	"server/middleware"
	"server/models"
	"server/services"
	"strings"
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
		Id:       primitive.NewObjectID(),
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

func (userController *UserController) Auth(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	headerAuthorization := c.GetReqHeaders()["Authorization"]
	token := strings.Split(headerAuthorization, " ")[1]
	if token == "" {
		return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateUserDTO{Success: false})
	}
	data := middleware.GetAccountId(token)
	accountIdStr, _ := data.(string)

	accountId, _ := primitive.ObjectIDFromHex(accountIdStr)

	account := userController.accountService.GetAccountById(ctx, accountId)
	if account == nil {
		return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	user := userController.userService.GetUserById(ctx, account.UserId)

	if user == nil {
		return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	return c.Status(http.StatusAccepted).JSON(dto.ResponseCreateUserDTO{Success: true, User: *user})
}
