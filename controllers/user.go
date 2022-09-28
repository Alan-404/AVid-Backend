package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
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

	var userData dto.CreateUserDTO

	if err := c.BodyParser(&userData); err != nil {
		fmt.Println(err)
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

	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false})
	}

	file, _ := fileHeader.Open()
	data, _ := ioutil.ReadAll(file)

	postFix := ".jpg"

	if strings.Split(http.DetectContentType(data), "/")[0] != "image" {
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false, Message: "Not Allow Anything else Image File"})
	}

	err = ioutil.WriteFile("./storage/avatar/"+user.Id.Hex()+postFix, data, 0644)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusBadRequest).JSON(dto.ResponseCreateUserDTO{Success: false})
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
	accountId, _ := primitive.ObjectIDFromHex(middleware.GetAccountId(token))

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

func (userController *UserController) GetAvatar(c *fiber.Ctx) error {
	id := c.Query("id")

	return c.Status(http.StatusAccepted).SendFile("./storage/avatar/" + id + ".jpg")
}
