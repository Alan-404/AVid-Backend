package dto

import (
	"os"
	"server/models"
)

type CreateUserDTO struct {
	FirstName string  `form:"firstName" bson:"firstName"`
	LastName  string  `form:"lastName" bson:"lastName"`
	Email     string  `form:"email"`
	Phone     string  `form:"phone"`
	Gender    string  `form:"gender"`
	BDate     string  `form:"bdate"`
	Password  string  `form:"password"`
	Role      bool    `form:"role"`
	Avatar    os.File `form:"avatar"`
}

type ResponseCreateUserDTO struct {
	Success bool        `json:"success"`
	User    models.User `json:"user"`
	Message string      `json:"message"`
}
