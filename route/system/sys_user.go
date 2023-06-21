package system

import (
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

func (u *SysUserRouter) AddSysUserRouter(route *gin.RouterGroup) {
	userRoute := route.Group("user")
	{
		userRoute.POST("", userApi.CreateUser)
		userRoute.GET("list", userApi.ListUser)
		userRoute.GET(":userid", userApi.GetUserById)
		userRoute.POST("update", userApi.UpdateUser)
		userRoute.DELETE(":userid", userApi.DeleteUser)
	}
}
