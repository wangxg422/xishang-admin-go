package public

import (
	"backend/common/response"
	"backend/initial/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginApi struct{}

func (m *LoginApi) Login(c *gin.Context) {
	response.Ok(c)
}

func (m *LoginApi) Logout(c *gin.Context) {
	logger.Info("注销成功")
	response.Ok(c)
}

func (m *LoginApi) GetInfo(c *gin.Context) {
	res := make(map[string]any)

	// roles, err := roleService.GetRolesByUserId(1)
	// if err != nil {
	// 	logger.Error("", zap.Error(err))
	// }

	userInfo, err := userService.GetUserInfo(1)
	if err != nil {
		logger.Error("", zap.Error(err))
	}

	res["permissions"] = []string{"*:*:*"}
	//res["roles"] = roles
	res["user"] = userInfo
	//res["roleIds"] = nil
	//res["deptIds"] = nil
	//res["postIds"] = nil
	res["admin"] = true

	logger.Info("注销成功")
	response.OkWithData(res, c)
}

func (m *LoginApi) Register(c *gin.Context) {
	response.Ok(c)
}
