package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ChannelId   primitive.ObjectID `json:"channelId" bson:"channelId"`
	Name        string             `json:"name" bson:"name"`
	Size        string             `json:"size" bson:"size"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	Description string             `json:"description" bson:"description"`
	View        int                `json:"view" bson:"view"`
}
