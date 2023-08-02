package system

import "github.com/gin-gonic/gin"

type SysPermissionRouter struct {
}

func (m *SysPermissionRouter) AddSysPermissionRouter(route *gin.RouterGroup) {
	router := route.Group("perm")
	{
		router.POST("", permissionApi.CreatePermission)
		//router.GET("", permissionApi.GetPermission)
		router.GET(":permId", permissionApi.GetPermissionById)
		router.PUT("", permissionApi.UpdatePermission)
		router.DELETE(":permId", permissionApi.DeletePermission)
	}
}
