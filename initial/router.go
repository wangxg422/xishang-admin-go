package initial

import (
	"backend/global"
	"backend/middleware"
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

	// 开放接口,健康检查、登录、注册、忘记密码、接口文档等
	publicGroup := r.Group("")
	{
		public.PublicRouterGroupApp.PublicRouter.AddPublicRouter(publicGroup)
	}

	authGroup := r.Group("")
	authGroup.Use(middleware.JWTAuth())

	sysGroup := authGroup.Group("sys")
	{
		system.SysRouterGroupApp.SysUserRouter.AddSysUserRouter(sysGroup)
		system.SysRouterGroupApp.SysDeptRouter.AddSysDeptRouter(sysGroup)
		system.SysRouterGroupApp.SysRoleRouter.AddSysRoleRouter(sysGroup)
		system.SysRouterGroupApp.SysPostRouter.AddSysPostRouter(sysGroup)
		system.SysRouterGroupApp.SysMenuRouter.AddSysMenuRouter(sysGroup)
		system.SysRouterGroupApp.SysDictRouter.AddSysDictRouter(sysGroup)
		system.SysRouterGroupApp.SysConfigRouter.AddSysConfigRouter(sysGroup)
	}

	_ = r.Run(":" + global.AppConfig.App.Port)
}
