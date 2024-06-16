package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// JWtCustomClaims represents the custom claims included in the JWT token.
type JWtCustomClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

// stSignKey is the secret signing key used to sign and verify JWT tokens.
var stSignKey = []byte(viper.GetString("jwt.signingKey"))

// GenerateToken generates a new JWT token with the provided user ID and name.
// It creates a JWT token with the custom claims `ID` and `Name`, as well as the registered claims `ExpiresAt`, `IssuedAt`, and `Subject`.
// The token is signed using the secret signing key `stSignKey`.
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

// ParseToken parses a JWT token and returns the custom claims.
// It takes a JWT token string as input and returns the `JWtCustomClaims` struct containing the custom claims, as well as any error encountered during the parsing process.
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

// IsTokenValid checks if a given JWT token is valid.
// It takes a JWT token string as input and returns a boolean indicating whether the token is valid or not.
func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	return err == nil
}
