package system

import "github.com/gin-gonic/gin"

type SysConfigRouter struct {
}

func (m *SysConfigRouter) AddSysConfigRouter(route *gin.RouterGroup) {
	router := route.Group("config")
	{
		//router.POST("", configApi.CreateConfig)
		//router.GET("list", configApi.ListConfig)
		router.GET("", configApi.GetConfig)
		//router.POST("update", configApi.UpdateConfig)
		//router.DELETE(":deptId", configApi.DeleteConfig)

		router.GET("configKey/:configKey", configApi.GetConfigByKey)
	}
}
