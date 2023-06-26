package system

import (
	"github.com/gin-gonic/gin"
)

type SysDeptRouter struct {
}

func (m *SysDeptRouter) AddSysDeptRouter(route *gin.RouterGroup) {
	router := route.Group("dept")
	{
		router.POST("", deptApi.CreateDept)
		router.GET("list", deptApi.ListDept)
		router.GET(":deptId", deptApi.GetDeptById)
		router.POST("update", deptApi.UpdateDept)
		router.DELETE(":deptId", deptApi.DeleteDept)
	}
}
