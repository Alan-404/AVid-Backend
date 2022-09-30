package services

import (
	"context"
	"server/configs"
	"server/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	accountCollection *mongo.Collection
}

func NewAccountService() *AccountService {
	accountService := new(AccountService)
	accountService.accountCollection = configs.GetCollection(configs.DB, "account")
	return accountService
}

func (accountService *AccountService) CreateAccount(ctx context.Context, account models.Account) *models.Account {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)

	account.Password = string(hashPassword)

	_, err := accountService.accountCollection.InsertOne(ctx, account)
	if err != nil {
		return nil
	}
	return &account

}

func (accountService *AccountService) GetAccountByUserId(ctx context.Context, userId primitive.ObjectID) *models.Account {
	var account *models.Account

	err := accountService.accountCollection.FindOne(ctx, &fiber.Map{"userId": userId}).Decode(&account)

	if err != nil {
		return nil
	}

	return account
}

func (accountService *AccountService) CheckPassword(account *models.Account, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))

	if err != nil {
		return false
	}

	return true
}

func (accountService *AccountService) GetAccountById(ctx context.Context, accountId primitive.ObjectID) *models.Account {
	var account *models.Account

	err := accountService.accountCollection.FindOne(ctx, &fiber.Map{"_id": accountId}).Decode(&account)

	if err != nil {
		return nil
	}

	return account

}

func (accountService *AccountService) ChangePassword(ctx context.Context, account *models.Account, oldPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(oldPassword))
	if err != nil {
		return false
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	update := &fiber.Map{"password": hashPassword}

	_, err = accountService.accountCollection.UpdateOne(ctx, &fiber.Map{"_id": account.Id}, &fiber.Map{"$set": update})

	if err != nil {
		return false
	}

	return true
}
