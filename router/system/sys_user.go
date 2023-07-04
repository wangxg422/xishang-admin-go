package system

import (
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {
}

func (m *SysUserRouter) AddSysUserRouter(route *gin.RouterGroup) {
	userRouter := route.Group("user")
	{
		userRouter.POST("", userApi.CreateUser)
		userRouter.GET("list", userApi.ListUserPage)
		userRouter.GET(":userId", userApi.GetUserById)
		userRouter.POST("update", userApi.UpdateUser)
		userRouter.DELETE(":userId", userApi.DeleteUser)
	}

	profileRoute := userRouter.Group("profile")
	{
		profileRoute.GET("profile", userApi.GetProfile)
	}
}
