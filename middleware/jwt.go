package middleware

import (
	"server/common"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateToken(accountId primitive.ObjectID) string {

	claims := jwt.MapClaims{
		"accountId": accountId.Hex(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(common.Secret))

	if err != nil {
		return ""
	}

	return accessToken

}

func GetAccountId(c *fiber.Ctx) error {
	headerAuthorization := c.GetReqHeaders()["Authorization"]
	token := strings.Split(headerAuthorization, " ")[1]

	if token == "" {
		return c.Status(400).JSON(&fiber.Map{})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Secret), nil
	})

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{})
	}

	for key, val := range claims {
		if key == "accountId" {
			accountId, _ := val.(string)
			c.Request().Header.Set("id", accountId)
			c.Next()
			return nil
		}
	}
	return c.Status(400).JSON(&fiber.Map{})
}
