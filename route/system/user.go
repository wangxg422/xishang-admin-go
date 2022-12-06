package system

import (
	sysApi "backend/api/v1/system"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
}

var userApi = sysApi.UserApi{}

func (s *UserRoute) AddSystemRoute(route *gin.RouterGroup) {
	systemRoute := route.Group("sys")
	{
		userRoute := systemRoute.Group("user")
		{
			userRoute.POST("", userApi.CreateUser)
			userRoute.GET("list", userApi.ListUser)
			userRoute.GET(":userid", userApi.GetUserById)
			userRoute.POST(":userid/update", userApi.UpdateUser)
			userRoute.DELETE(":userid", userApi.DeleteUser)
		}
	}
}
