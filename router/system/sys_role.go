package system

import (
	"github.com/gin-gonic/gin"
)

type SysRoleRouter struct {
}

func (m *SysRoleRouter) AddSysRoleRouter(route *gin.RouterGroup) {
	router := route.Group("role")
	{
		router.POST("", roleApi.CreateRole)
		router.GET("list", roleApi.ListRole)
		router.GET(":roleid", roleApi.GetRoleById)
		router.POST("update", roleApi.UpdateRole)
		router.DELETE(":roleid", roleApi.DeleteRole)
	}
}
