package system

import (
	"backend/common/response"
	"backend/global"
	"backend/initial/logger"
	"backend/model/system"
	"backend/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

type JwtService struct{}

func (m *JwtService) IsBlacklist(token string) bool {
	return false
}

func (m *JwtService) GetRedisJWT(username string) (string, error) {
	redisJWT, err := global.RedisClient.Get(context.Background(), username).Result()
	return redisJWT, err
}

func (m *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.AppConfig.Jwt.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.RedisClient.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func (m *JwtService) SignToken(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.AppConfig.Jwt.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		logger.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.AppConfig.App.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := m.GetRedisJWT(user.Username); err == redis.Nil {
		if err := m.SetRedisJWT(token, user.Username); err != nil {
			logger.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		logger.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := m.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := m.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
// 	err = global.DB.Create(&jwtList).Error
// 	if err != nil {
// 		return
// 	}
// 	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
// 	return
// }
