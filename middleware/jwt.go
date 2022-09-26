package middleware

import (
	"server/common"

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

func GetAccountId(accessToken string) interface{} {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(common.Secret), nil
	})

	if err != nil {
		return nil
	}

	for key, val := range claims {
		if key == "accountId" {
			return val
		}
	}
	return nil
}
