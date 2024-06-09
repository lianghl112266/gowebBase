package router

import (
	"context"
	"fmt"
	_ "goweb/docs"
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

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func initBasePlatformRoutes() {
	InitUserRoutes()
}

func InitRouter() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.Use(middleware.Cors())

	rgPublic := r.Group("./api/v1/public")
	rgAuth := r.Group("./api/vi")
	rgAuth.Use(middleware.Auth())
	//注册基本平台基本路由
	initBasePlatformRoutes()
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	//注册验证器
	registerCustomValidator()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}
	//err := r.Run(fmt.Sprintf(":%s", stPort))
	//if err != nil {
	//	panic(fmt.Sprintf("Start Server Error: %s", err.Error()))
	//}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", stPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()

	//graceful shutdown
	<-ctx.Done()

	global.Logger.Info(fmt.Sprintf("Stop Listen: %s", stPort))
}

func registerCustomValidator() {
	if engine, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = engine.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			value, ok := fl.Field().Interface().(string)
			return ok && len(value) > 0 && value[0] == 'a'
		})
	}
}
