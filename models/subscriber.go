package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscriber struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ChannelId   primitive.ObjectID `json:"channelId" bson:"channelid"`
	UserId      primitive.ObjectID `json:"userId" bson:"userid"`
	SubscibedAt time.Time          `json:"subscribedAt" bson:"subscribedat"`
	Status      bool               `json:"status" bson:"status"`
}
