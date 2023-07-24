package jwt

import (
	"backend/common/constant"
	"backend/initial/logger"
	"backend/model/common/request"
	"backend/utils"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*request.CustomClaims, error) {
	token := c.Request.Header.Get(constant.TokenName)
	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		logger.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) int64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.UserId
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.BaseClaims.UserId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *request.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.Username
	}
}
