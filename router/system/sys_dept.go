package system

import (
	"github.com/gin-gonic/gin"
)

type SysDeptRouter struct {
}

func (s *SysDeptRouter) AddSysDeptRouter(route *gin.RouterGroup) {
	router := route.Group("dept")
	{
		router.POST("", deptApi.CreateDept)
		router.GET("list", deptApi.ListDept)
		router.GET(":deptid", deptApi.GetDeptById)
		router.POST("update", deptApi.UpdateDept)
		router.DELETE(":deptid", deptApi.DeleteDept)
	}
}
