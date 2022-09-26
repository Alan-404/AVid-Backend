package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	UserId   primitive.ObjectID `json:"userId"`
	Password string             `json:"password"`
	Role     bool               `json:"role"`
}
