package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"goweb/conf"
)

var (
	Logger      *zap.SugaredLogger
	DB          *gorm.DB
	RedisClient *conf.RedisClient
)
