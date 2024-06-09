package router

import (
	"github.com/gin-gonic/gin"
	"goweb/api"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			rgAuthUser.POST("/list", userApi.GetUserList)
		}
	})
}
