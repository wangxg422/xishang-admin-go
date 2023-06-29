package global

import (
	"backend/config"

	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                      *gorm.DB
	AppConfig               config.AppConfig
	Log                     *zap.Logger
	RedisClient             *redis.Client
	GVA_Concurrency_Control = &singleflight.Group{}
	LocalCache              local_cache.Cache
)
