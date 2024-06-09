package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

type JWtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

var stSignKey = []byte(viper.GetString("jwt.signingKey"))

func GenerateToken(id uint, name string) (string, error) {
	jwtCustomClaims := JWtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCustomClaims)
	return token.SignedString(stSignKey)
}

func ParseToken(token string) (JWtCustomClaims, error) {
	jwtCustomClaims := JWtCustomClaims{}
	parseToken, err := jwt.ParseWithClaims(token, &jwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})
	if err == nil && !parseToken.Valid {
		err = errors.New("invalid Token")
	}
	return jwtCustomClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
