// Package middleware provides middleware functions for Gin framework.
//
// Middleware functions are executed before the actual handler function
// of a route. They allow you to perform actions like CORS handling,
// authentication, authorization, logging, and more, before the request
// is processed.
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors is a middleware function that enables Cross-Origin Resource Sharing (CORS)
// for your API.
//
// It allows requests from different origins (domains, protocols, ports)
// to access your API.
func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		//AllowOrigins: []string{}, // You can specify allowed origins here
		//AllowAllOrigins: true, // Or allow all origins
		AllowOriginFunc: func(_ string) bool {
			return true // Allow all origins
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Context-Type", "Authorization", "Accept", "token"},
		AllowCredentials: true, // Allow credentials to be sent with requests
	}
	return cors.New(cfg)
}
