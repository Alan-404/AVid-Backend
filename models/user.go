package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email"`
	Phone     string             `json:"phone"`
	Gender    string             `json:"gender"`
	BDate     string             `json:"bDate" bson:"bDate"`
}
