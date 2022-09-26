package dto

import "server/models"

type CreateUserDTO struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	BDate     string `json:"bdate"`
	Password  string `json:"password"`
	Role      bool   `json:"role"`
}

type ResponseCreateUserDTO struct {
	Success bool        `json:"success"`
	User    models.User `json:"user"`
	Message string      `json:"message"`
}
