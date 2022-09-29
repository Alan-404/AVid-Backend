package controllers

import (
	"context"
	"server/models"
	"server/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChannelController struct {
	userService    *services.UserService
	accountService *services.AccountService
	channelService *services.ChannelService
}

func NewChannelController() *ChannelController {
	channelController := new(ChannelController)
	channelController.accountService = services.NewAccountService()
	channelController.userService = services.NewUserService()
	channelController.channelService = services.NewChannelService()
	return channelController
}

func (channelController *ChannelController) CreateChannel(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	accountId, _ := primitive.ObjectIDFromHex(c.GetReqHeaders()["Id"])

	account := channelController.accountService.GetAccountById(ctx, accountId)
	if account == nil {
		return c.Status(400).JSON(&fiber.Map{})
	}

	var channel = *&models.Channel{
		Id:        primitive.NewObjectID(),
		UserId:    account.UserId,
		Subscribe: 0,
		CreatedAt: time.Now(),
		Status:    true,
	}

	addedChannel := channelController.channelService.CreateChannel(ctx, channel)

	if addedChannel == nil {
		return c.Status(400).JSON(&fiber.Map{})
	}

	return c.Status(200).JSON(&fiber.Map{"channel": channel})
}
