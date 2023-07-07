package system

import "github.com/gin-gonic/gin"

type SysConfigRouter struct {
}

func (m *SysConfigRouter) AddSysConfigRouter(route *gin.RouterGroup) {
	router := route.Group("config")
	{
		router.POST("", configApi.CreateConfig)
		router.GET("list", configApi.GetConfigPage)
		router.GET(":configId", configApi.GetConfigById)
		router.PUT("", configApi.UpdateConfig)
		router.DELETE(":configId", configApi.DeleteConfig)

		router.GET("configKey/:configKey", configApi.GetConfigByKey)
	}
}
