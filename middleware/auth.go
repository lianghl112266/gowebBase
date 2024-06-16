// Package middleware provides middleware functions for Gin framework.
//
// Middleware functions are executed before the actual handler function
// of a route. They allow you to perform actions like authentication,
// logging, cors, and more, before the request is processed.
package middleware

import (
	"goweb/api"
	"goweb/global"
	"goweb/model"
	"goweb/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ErrCodeInvalidToken is a custom error code for invalid tokens.
const ErrCodeInvalidToken = 10005

// Auth is a middleware function for authentication using JWT tokens.
// It verifies the token, checks its validity, and sets the user information
// in the context.
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		// If the request doesn't have a token, return an unauthorized error.
		if token == "" {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		// Parse the token and check for errors.
		jwtCustomClaim, err := utils.ParseToken(token)
		if err != nil {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		// Check if the token matches the one stored in Redis.
		id := strconv.Itoa(int(jwtCustomClaim.ID))
		redisToken, err := global.RedisClient.Get(id)
		if err != nil || redisToken != token {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		// Check if the token is expired.
		d, err := global.RedisClient.GetExpireDuration(id)
		if err != nil || d <= 0 {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		// Proceed to the next handler in the chain.
		defer ctx.Next()

		// Renew the token in Redis if it's about to expire.
		if d <= 10*time.Minute {
			_ = global.RedisClient.Set(id, token, 20*time.Minute)
			//_ = global.RedisClient.Set(id, token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
		}

		global.Logger.Info("id: ", id, ",name: ", jwtCustomClaim.Name, ",url: ", ctx.Request.RequestURI)

		// Set the user information in the context.
		ctx.Set("LoginUser", model.LoginUser{
			ID:   id,
			Name: jwtCustomClaim.Name,
		})
	}
}
