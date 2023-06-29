package system

import (
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

func (m *SysUserRouter) AddSysUserRouter(route *gin.RouterGroup) {
	router := route.Group("user")
	{
		router.POST("", userApi.CreateUser)
		router.GET("list", userApi.ListUser)
		router.GET(":userId", userApi.GetUserById)
		router.POST("update", userApi.UpdateUser)
		router.DELETE(":userId", userApi.DeleteUser)
		
		route.GET("profile", userApi.GetProfile)
	}
}
