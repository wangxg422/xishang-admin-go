package public

import (
	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

func (m *PublicRouter) AddPublicRouter(route *gin.RouterGroup) {
	// 健康检查
	route.GET("health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 登录、注销、注册
	route.POST("login", loginApi.Login)
	route.POST("logout", loginApi.Logout)
	route.POST("register", loginApi.Register)

	// 获取用户信息
	route.GET("getInfo", loginApi.GetInfo)

	// 获取验证码
	route.GET("captcha", captchaApi.GenCaptcha)
}
