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
		menuRouter.GET("list", menuApi.ListMenu)
		menuRouter.GET(":menuId", menuApi.GetMenuById)
		menuRouter.POST("update", menuApi.UpdateMenu)
		menuRouter.DELETE(":menuId", menuApi.DeleteMenu)
	}

	dynamicRouter := router.Group("router")
	{
		dynamicRouter.GET("", menuApi.GetMenuByUser)
	}
}
