package cmd

import (
	"fmt"
	"goweb/conf"
	"goweb/global"
	"goweb/router"
)

// Start function start the application by setting up configurations, logging, database, and Redis connections.
func Start() {
	// Initialize application configurations
	err := conf.InitConfig()
	if err != nil {
		// Log an error if configuration initialization fails
		global.Logger.Error("Failed to initialize configuration: %v", err.Error())
		return
	}

	// Initialize the logger using the configuration settings
	global.Logger = conf.InitLogger()

	// Initialize the database connection
	db, err := conf.InitDB()
	if err != nil {
		// Log an error if database initialization fails
		global.Logger.Error(fmt.Sprintf("Failed to initialize DB: %v", err.Error()))
		return
	} else {
		// Print a success message if database initialization is successful
		fmt.Println("db init successful")
	}
	// Set the global database variable
	global.DB = db

	// Initialize the Redis client
	rdClient, err := conf.InitRedis()
	if err != nil {
		// Log an error if Redis initialization fails
		global.Logger.Error(fmt.Sprintf("Failed to initialize Redis: %v", err.Error()))
		return
	}
	// Set the global Redis client variable
	global.RedisClient = rdClient

	// Initialize the application router
	router.InitRouter()
}

// Clean function is likely used for cleanup operations before the application exits,
// such as closing database connections or releasing resources.
func Clean() {
	fmt.Println("-----------------Clean------------")
}
