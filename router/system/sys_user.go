package system

import (
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

func (u *SysUserRouter) AddSysUserRouter(route *gin.RouterGroup) {
	router := route.Group("user")
	{
		router.POST("", userApi.CreateUser)
		router.GET("list", userApi.ListUser)
		router.GET(":userid", userApi.GetUserById)
		router.POST("update", userApi.UpdateUser)
		router.DELETE(":userid", userApi.DeleteUser)
	}
}
