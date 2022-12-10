package initial

import (
	"backend/global"
	"backend/route"
	"github.com/gin-gonic/gin"
)

func InitRoute() {
	// debug or release
	gin.SetMode(global.APP_CONFIG.App.Mode)
	r := gin.Default()

	// 开放接口，健康检查、注册、忘记密码等
	publicRouteGroup := route.AppRouteGroupIns.PublicRouteGroup
	publicGroup := r.Group("public")
	{
		publicRouteGroup.AddPublicGroup(publicGroup)
	}

	sysRouteGroup := route.AppRouteGroupIns.SysRouteGroup
	privateGroup := r.Group("")
	{
		sysRouteGroup.AddSystemRoute(privateGroup)
	}

	_ = r.Run(":8081")
}
