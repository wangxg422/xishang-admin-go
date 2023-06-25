package initial

import (
	"backend/global"
	"backend/router/public"
	"backend/router/system"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// debug or release
	gin.SetMode(global.AppConfig.App.Mode)
	r := gin.Default()

	if addr := global.AppConfig.App.Addresses; addr == "" || addr == "0.0.0.0" {

	} else {
		r.SetTrustedProxies(strings.Split(global.AppConfig.App.Addresses, ","))
	}

	// 开放接口,健康检查、注册、忘记密码、接口文档等
	publicRouterGroup := &public.PublicRouterGroup{}
	publicGroup := r.Group("public")
	{
		publicRouterGroup.AddPublicRouterGroup(publicGroup)
	}

	sysGroup := r.Group("sys")
	{
		system.SysRouterGroupApp.SysUserRouter.AddSysUserRouter(sysGroup)
		system.SysRouterGroupApp.SysDeptRouter.AddSysDeptRouter(sysGroup)
	}

	_ = r.Run(":" + global.AppConfig.App.Port)
}
