package services

import (
	"context"
	"server/configs"
	"server/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberService struct {
	subscriberCollection *mongo.Collection
}

func NewSubscriberService() *SubscriberService {
	subscriberService := new(SubscriberService)

	subscriberService.subscriberCollection = configs.GetCollection(configs.DB, "subscriber")

	return subscriberService
}

func (subscriberService *SubscriberService) AddSubscriber(ctx context.Context, subscriber models.Subscriber) *models.Subscriber {

	_, err := subscriberService.subscriberCollection.InsertOne(ctx, subscriber)

	if err != nil {
		return nil
	}

	return &subscriber

}
