package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscriber struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ChannelId   primitive.ObjectID `json:"channelId" bson:"channelId"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	SubscibedAt time.Time          `json:"subscribedAt" bson:"subscribedAt"`
	Status      bool               `json:"status" bson:"status"`
}
