package token

import (
	"github.com/dgrijalva/jwt-go"
	"uangq-be/common/logger"
	"uangq-be/configs"
)

func GenerateJWT(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	tokenString, err := token.SignedString(configs.SessionJwtSecret)
	if err != nil {
		logger.Log("Fail to create login token. " + err.Error())
		return ""
	}

	return tokenString
}
