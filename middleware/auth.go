package middleware

import (
	"github.com/gin-gonic/gin"
	"goweb/api"
	"goweb/global"
	"goweb/model"
	"goweb/utils"
	"net/http"
	"strconv"
	"time"
)

const ErrCodeInvalidToken = 10005

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		//请求时没有携带token
		if token == "" {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		//token错误
		jwtCustomClaim, err := utils.ParseToken(token)
		if err != nil {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}

		//不等于redis的token
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

		//token过期
		d, err := global.RedisClient.GetExpireDuration(id)
		if err != nil || d <= 0 {
			api.Fail(ctx, api.ResponseJson{
				Status: http.StatusUnauthorized,
				Code:   ErrCodeInvalidToken,
				Msg:    "invalid token",
			})
			return
		}
		defer ctx.Next()

		//token续期
		if d <= 10*time.Minute {
			_ = global.RedisClient.Set(id, token, 20*time.Minute)
			//_ = global.RedisClient.Set(id, token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
		}

		ctx.Set("LoginUser", model.LoginUser{
			ID:   id,
			Name: jwtCustomClaim.Name,
		})
	}
}
