package system

import "github.com/gin-gonic/gin"

type SysModuleRouter struct {
}

func (m *SysModuleRouter) AddSysModuleRouter(route *gin.RouterGroup) {
	router := route.Group("module")
	{
		router.POST("", moduleApi.CreateModule)
		router.GET("", moduleApi.GetModule)
		router.GET(":moduleId", moduleApi.GetModuleById)
		router.PUT("", moduleApi.UpdateModule)
		router.DELETE(":moduleId", moduleApi.DeleteModule)
	}
}
