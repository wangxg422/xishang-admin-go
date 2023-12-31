package system

import "github.com/gin-gonic/gin"

type SysDictRouter struct {
}

func (m *SysDictRouter) AddSysDictRouter(route *gin.RouterGroup) {
	router := route.Group("dict")

	dataRouter := router.Group("data")
	{
		dataRouter.GET("type/:dictType", dictApi.GetDictDataByType)
		dataRouter.GET("list", dictApi.GetDictDataPage)
		dataRouter.POST("", dictApi.CreateDictData)
		dataRouter.PUT("", dictApi.UpdateDictData)
		dataRouter.GET(":dictDataId", dictApi.GetDictDataById)
		dataRouter.DELETE(":dictDataId", dictApi.DeleteDictData)
	}

	typeRouter := router.Group("type")
	{
		typeRouter.GET("list", dictApi.GetDictTypePage)
		typeRouter.POST("", dictApi.CreateDictType)
		typeRouter.GET(":dictTypeId", dictApi.GetDictTypeById)
		typeRouter.PUT("", dictApi.UpdateDictType)
		typeRouter.DELETE(":dictTypeId", dictApi.DeleteDictType)
		typeRouter.GET("list/all", dictApi.GetDictTypeAll)
	}
}
