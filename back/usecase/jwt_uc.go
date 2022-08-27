package usecase

import (
	"ideal-journey/clients/errors"
	"ideal-journey/clients/logger"
	"ideal-journey/entity"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtUC struct {
}

func NewJwtUC() *JwtUC {
	return &JwtUC{}
}

func (s *JwtUC) GenerateJWT(user *entity.UserAuth) (string, errors.RestErr) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = user.Id
	if strings.TrimSpace(user.Fingerprint) == "" {
		claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	} else {
		claims["fingerprint"] = user.Fingerprint
	}

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		logger.Error("[JWT] Error", err)
		return "", errors.InternalServerError("jwt_error", err)
	}
	return tokenString, nil
}
