package cmd

import (
	"fmt"
	"goweb/conf"
	"goweb/global"
	"goweb/router"
)

func Start() {
	err := conf.InitConfig()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Conf Init Error: %v", err.Error()))
	}

	global.Logger = conf.InitLogger()

	db, err := conf.InitDB()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("DB Init Error: %v", err.Error()))
		return
	} else {
		fmt.Println("db init successful")
	}
	global.DB = db

	rdClient, err := conf.InitRedis()
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Redis Init Error: %v", err.Error()))
		return
	}
	global.RedisClient = rdClient

	router.InitRouter()
}

func Clean() {
	fmt.Println("-----------------Clean------------")
}
