package permission

import (
	"backend/global"
	"backend/initial/logger"
	"backend/model/common/response"
	"backend/utils/jwt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	dbAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

const basicModel string = `
[request_definition]
r = sub, mod, obj, act

[policy_definition]
p = sub, mod, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.obj == p.obj && g(r.sub, p.sub) && r.act == p.act || r.sub == "admin"
`

func InitAdapter() (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(basicModel)
	if err != nil {
		zap.L().Error("连接数据库失败", zap.Error(err))
		return nil, err
	}

	a, err := dbAdapter.NewAdapterByDB(global.DB)
	if err != nil {
		return nil, err
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return e, err
	}

	return e, nil
}

func CasbinPermCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.AppConfig.App.Mode == "debug" {
			userId := jwt.GetUserID(c)
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, "")
			// 获取请求方法
			//act := c.Request.Method

			sub := strconv.FormatInt(userId, 10)
			e, err := InitAdapter()
			if err != nil {
				logger.Error("", zap.Error(err))
			}

			success, err := e.Enforce(sub, obj, nil)
			if err != nil {
				logger.Error("", zap.Error(err))
				response.FailWithMessage(err.Error(), c)
				c.Abort()
				return
			}
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
