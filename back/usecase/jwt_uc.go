package usecase

import (
	"ideal-journey/clients/errors"
	"ideal-journey/clients/logger"
	"ideal-journey/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtUC struct {
}

func NewJwtUC() *JwtUC {
	return &JwtUC{}
}

type MyCustomClaims struct {
	Subject     string `json:"Subject"`
	ExpiresAt   int64  `json:"ExpiresAt"`
	Fingerprint string `json:"Fingerprint"`
	jwt.RegisteredClaims
}

func (s *JwtUC) ValidateJWT(tokenString, fingerprint string) errors.RestErr {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	if err != nil || !token.Valid {
		return errors.UnautorizedError()
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		if strings.TrimSpace(claims.Fingerprint) == "" && IsExpired(claims.ExpiresAt) {
			return errors.UnautorizedError()
		} else if strings.TrimSpace(claims.Fingerprint) != "" {
			return validateFingerprint(claims.Fingerprint, fingerprint)
		}
	}
	return nil
}

func validateFingerprint(tokenFingerprint, fingerprint string) errors.RestErr {
	if tokenFingerprint != fingerprint {
		return errors.UnautorizedError()
	}
	return nil
}

func IsExpired(expireAt int64) bool {
	return time.Unix(expireAt, 0).Before(time.Now().UTC())
}

func (s *JwtUC) GenerateJWT(user *entity.UserAuth) (string, errors.RestErr) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["Subject"] = user.Id
	if strings.TrimSpace(user.Fingerprint) == "" {
		claims["ExpiresAt"] = time.Now().Add(time.Hour * 3).Unix()
	} else {
		claims["Fingerprint"] = user.Fingerprint
	}

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		logger.Error("[JWT] Error", err)
		return "", errors.InternalServerError("jwt_error", err)
	}
	return tokenString, nil
}
