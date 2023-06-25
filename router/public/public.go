package public

import (
	"backend/api/v1/public"

	"github.com/gin-gonic/gin"
)

type PublicRouter struct {
}

var loginAPi = public.LoginApi{}

func (m *PublicRouter) AddPublicRouter(route *gin.RouterGroup) {
	// 健康检查
	route.GET("health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 登录注册
	route.POST("login", loginAPi.Login)
	route.POST("register", loginAPi.Register)
}
