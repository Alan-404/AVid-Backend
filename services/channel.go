package services

import (
	"context"
	"server/configs"
	"server/models"

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
