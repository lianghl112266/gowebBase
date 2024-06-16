package router

import (
	"goweb/api"

	"github.com/gin-gonic/gin"
)

// InitUserRoutes initializes the routes related to user management.
func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		// Create a new instance of the user API.
		userApi := api.NewUserApi()

		// Define a public route group for user-related endpoints that don't require authentication.
		rgPublicUser := rgPublic.Group("user")
		{
			// Register the user login route.
			rgPublicUser.POST("/login", userApi.Login)
		}

		// Define an authenticated route group for user-related endpoints that require authentication.
		rgAuthUser := rgAuth.Group("user")
		{
			// Register the user creation route.
			rgAuthUser.POST("", userApi.AddUser)

			// Register the route to retrieve a user by ID.
			rgAuthUser.GET("/:id", userApi.GetUserById)

			// Register the route to retrieve a list of users.
			rgAuthUser.POST("/list", userApi.GetUserList)
		}
	})
}
