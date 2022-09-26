package services

import (
	"context"
	"server/configs"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userCollection *mongo.Collection
}

func NewUserService() *UserService {
	userService := new(UserService)
	userService.userCollection = configs.GetCollection(configs.DB, "user")

	return userService
}

func (userService *UserService) CreateUser(ctx context.Context, user models.User) *primitive.ObjectID {

	result, err := userService.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return &objId

}

func (userService *UserService) GetUserByEmail(ctx context.Context, email string) *models.User {
	var user *models.User

	err := userService.userCollection.FindOne(ctx, &fiber.Map{"email": email}).Decode(&user)

	if err != nil {
		return nil
	}

	return user
}

func (userService *UserService) GetUserById(ctx context.Context, id primitive.ObjectID) *models.User {
	var user *models.User

	err := userService.userCollection.FindOne(ctx, &fiber.Map{"_id": id}).Decode(&user)

	if err != nil {
		return nil
	}

	return user

}
