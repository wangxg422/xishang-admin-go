package global

import (
	"backend/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Viper       *viper.Viper
	AppConfig   config.AppConfig
	Log         *zap.Logger
	RedisClient *redis.Client
)
