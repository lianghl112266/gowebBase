/*
Package router defines the routing and initialization logic for the web application.

It manages the registration of different route groups, including public and authenticated routes.
It also handles the setup of the Gin web framework, middleware, and custom validators.

The package utilizes Swagger to generate API documentation, making it easier to understand
the available endpoints and their functionality.
*/

package router

import (
	"context"
	"fmt"
	_ "goweb/docs" // Import swagger documentation
	"goweb/global"
	"goweb/middleware"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Define the type for route registration functions
type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

// Global list of route registration functions
var (
	gfnRoutes []IFnRegisterRoute
)

// Register a route registration function
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// Initialize base platform routes
func initBasePlatformRoutes() {
	InitUserRoutes()
}

// Initialize the router
func InitRouter() {
	// Create a context with signal notification
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.Use(middleware.Cors())

	// Create a public route group for routes that don't require authentication
	rgPublic := r.Group("./api/v1/public")

	// Create an authenticated route group for routes that require authentication
	rgAuth := r.Group("./api/vi")
	rgAuth.Use(middleware.Auth())

	// RRegister all registered route functions
	initBasePlatformRoutes()
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// Register custom validator
	registerCustomValidator()

	// Register Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Get port number from configuration file, use default port 8999 if not found
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	// Create HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// Start the server
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", stPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()

	// Shutdown the server
	<-ctx.Done()

	global.Logger.Info(fmt.Sprintf("Stop Listen: %s", stPort))
}

// Register custom validator
func registerCustomValidator() {
	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Register custom validation rule: First character must be "a"
		_ = engine.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			value, ok := fl.Field().Interface().(string)
			return ok && len(value) > 0 && value[0] == 'a'
		})
	}
}
