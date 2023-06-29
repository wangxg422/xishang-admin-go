package public

import (
	"backend/common/response"
	"backend/global"
	"backend/initial/logger"
	"backend/model/system"
	"backend/utils/captcha"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginApi struct{}

func (m *LoginApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	captchaEnabled := global.AppConfig.Captcha.Enabled           // 是否开启验证码

	if captchaEnabled && captcha.GetRedisStore().Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			logger.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			logger.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		jwtService.SignToken(c,*user)
		return
	} else {
		response.FailWithMessage("验证码错误", c)
	}
	
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
