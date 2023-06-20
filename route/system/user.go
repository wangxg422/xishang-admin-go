package system

import (
	sysApi "backend/api/v1/system"

	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

var userApi = sysApi.UserApi{}

func (s *SysUserRouter) AddSysUserRouter(route *gin.RouterGroup) {
	userRoute := route.Group("user")
	{
		userRoute.POST("", userApi.CreateUser)
		userRoute.GET("list", userApi.ListUser)
		userRoute.GET(":userid", userApi.GetUserById)
		userRoute.POST("update", userApi.UpdateUser)
		userRoute.DELETE(":userid", userApi.DeleteUser)
	}
}
