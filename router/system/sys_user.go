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
		userRouter.GET("list", userApi.GetUserPage)
		userRouter.GET(":userId", userApi.GetUserById)
		userRouter.PUT("", userApi.UpdateUser)
		userRouter.DELETE(":userId", userApi.DeleteUser)
		userRouter.PUT("status", userApi.ChangeUserStatus)
		userRouter.PUT("password/reset", userApi.ChangeUserPassword)
	}

	profileRoute := userRouter.Group("profile")
	{
		profileRoute.GET("profile", userApi.GetProfile)
	}
}
