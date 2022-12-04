package system

import (
	v1 "backend/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
}

func (s *UserRoute) AddSystemRoute(route *gin.RouterGroup) {
	systemRoute := route.Group("sys")
	{
		userRoute := systemRoute.Group("user")
		userApi := v1.AppApiGroupIns.SysApiGroup.UserApi
		{
			userRoute.POST("", userApi.CreateUser)
			userRoute.GET("list", userApi.ListUser)
			userRoute.GET(":userid", userApi.GetUserById)
			//userRoute.POST("update", userApi.UpdateUser)
			//userRoute.DELETE(":id", userApi.DeleteUser)
		}
	}
}
