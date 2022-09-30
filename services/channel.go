package services

import (
	"context"
	"fmt"
	"server/configs"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChannelService struct {
	channelCollection *mongo.Collection
}

func NewChannelService() *ChannelService {
	channelService := new(ChannelService)

	channelService.channelCollection = configs.GetCollection(configs.DB, "channel")

	return channelService
}

func (channelService *ChannelService) CreateChannel(ctx context.Context, channel models.Channel) *models.Channel {
	_, err := channelService.channelCollection.InsertOne(ctx, channel)
	if err != nil {
		return nil
	}

	return &channel
}

func (channelService *ChannelService) GetChannelByUserId(ctx context.Context, userId primitive.ObjectID) *models.Channel {
	var channel *models.Channel

	err := channelService.channelCollection.FindOne(ctx, &fiber.Map{"userId": userId}).Decode(&channel)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return channel
}

func (channelService *ChannelService) GetChannelById(ctx context.Context, id primitive.ObjectID) *models.Channel {
	var channel models.Channel

	err := channelService.channelCollection.FindOne(ctx, &fiber.Map{"_id": id}).Decode(&channel)

	if err != nil {
		return nil
	}

	return &channel
}
