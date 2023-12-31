package public

import (
	"backend/common/enmu"
	"backend/global"
	"backend/initial/logger"
	"backend/model/common/response"
	"backend/model/dto/system"
	"backend/utils/captcha"

	set "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginApi struct{}

func (m *LoginApi) Login(c *gin.Context) {
	loginDto := system.SysLoginDTO{}
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	captchaEnabled := global.AppConfig.Captcha.Enabled

	// 检查验证码
	if captchaEnabled {
		// 检查验证码是否正确
		captchaMatch := captcha.GetRedisStore().Verify(loginDto.CaptchaId, loginDto.CaptchaCode, true)
		if !captchaMatch {
			response.FailWithMessage("验证码错误", c)
			return
		}
	}

	user, err := userService.GetUserByUserName(loginDto.UserName)
	if err != nil {
		logger.Error("登陆失败! 用户名密码错误!", zap.Error(err))
		response.FailWithMessage("用户名密码错误", c)
		return
	}
	if enmu.StatusDisabled.Equals(user.Status) {
		logger.Error("登陆失败! 用户被禁止登录!")
		response.FailWithMessage("用户被禁止登录", c)
		return
	}

	jwtService.SignToken(c, user)
}

func (m *LoginApi) Logout(c *gin.Context) {
	logger.Info("注销成功")
	response.Ok(c)
}

func (m *LoginApi) GetInfo(c *gin.Context) {
	res := make(map[string]any)

	userInfo, err := userService.GetUserInfo(1)
	if err != nil {
		logger.Error("", zap.Error(err))
	}

	roleSet := set.NewSet()
	if len(userInfo.SysRoles) > 0 {
		for _, r := range userInfo.SysRoles {
			roleSet.Add(r.RoleCode)
		}
	}

	res["permissions"] = []string{"*:*:*"}
	res["roles"] = roleSet
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
