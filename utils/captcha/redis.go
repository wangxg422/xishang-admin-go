package captcha

import (
	"backend/global"
	"backend/initial/logger"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

/**
单例模式
*/

var redisStoreInstance *RedisStore

var once sync.Once

func GetRedisStore() *RedisStore {
	once.Do(func() {
		redisStoreInstance = &RedisStore{
			Expiration: time.Duration(global.AppConfig.Captcha.Expiration * int64(time.Second)),
			PreKey:     global.AppConfig.Captcha.PreKey,
		}
	})

	return redisStoreInstance
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.RedisClient.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		logger.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.RedisClient.Get(rs.Context, key).Result()
	if err != nil {
		logger.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.RedisClient.Del(rs.Context, key).Err()
		if err != nil {
			logger.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	fmt.Println(key)
	v := rs.Get(key, clear)
	return v == answer
}
