package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	Size        string             `json:"size" bson:"size"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	Description string             `json:"description" bson:"description"`
	View        int                `json:"view" bson:"view"`
}
