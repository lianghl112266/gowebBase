package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		//AllowOrigins: []string{},
		//AllowAllOrigins: true,
		AllowOriginFunc: func(_ string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Context-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true,
	}
	return cors.New(cfg)
}
