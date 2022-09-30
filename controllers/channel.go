package controllers

import (
	"context"
	"server/dto"
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
		return c.Status(400).JSON(dto.ResponseCreateChannel{Success: false, Message: "Undentified Token"})
	}

	checkChannel := channelController.channelService.GetChannelByUserId(ctx, account.UserId)

	if checkChannel != nil {
		return c.Status(500).JSON(dto.ResponseCreateChannel{Success: false, Message: "Internel Error Server"})
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
		return c.Status(500).JSON(dto.ResponseCreateChannel{Success: false, Message: "Internel Error Server"})
	}

	return c.Status(200).JSON(dto.ResponseCreateChannel{Success: true, Channel: &channel})
}

func (channelControler *ChannelController) GetChannelById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(dto.ResponseGetChannelById{Success: false})
	}
	channelId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(500).JSON(dto.ResponseGetChannelById{Success: false})
	}
	channel := channelControler.channelService.GetChannelById(ctx, channelId)
	return c.Status(200).JSON(dto.ResponseGetChannelById{Success: true, Channel: channel})

}
