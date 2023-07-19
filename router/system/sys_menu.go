package system

import (
	"github.com/gin-gonic/gin"
)

type SysMenuRouter struct {
}

func (m *SysMenuRouter) AddSysMenuRouter(router *gin.RouterGroup) {
	menuRouter := router.Group("menu")
	{
		menuRouter.POST("", menuApi.CreateMenu)
		menuRouter.GET("list", menuApi.GetMenu)
		menuRouter.GET(":menuId", menuApi.GetMenuById)
		menuRouter.PUT("", menuApi.UpdateMenu)
		menuRouter.DELETE(":menuId", menuApi.DeleteMenu)

		menuRouter.GET("tree", menuApi.GetMenuTree)
	}

	dynamicRouter := router.Group("router")
	{
		dynamicRouter.GET("", menuApi.GetRouterByUserId)
	}
}
