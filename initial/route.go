package initial

import (
	"backend/route"
	"github.com/gin-gonic/gin"
)

func InitRoute() {

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
