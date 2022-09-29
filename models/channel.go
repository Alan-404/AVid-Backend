package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Channel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Subscribe int                `json:"subscribe" bson:"subscibe"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdat"`
	Status    bool               `json:"status" bson:"status"`
}
