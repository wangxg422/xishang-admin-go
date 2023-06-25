package system

import (
	"backend/common/enmu"
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDeptApi struct{}

func (d *SysDeptApi) CreateDept(c *gin.Context) {
	deptDto := dto.SysCreateDeptDTO{}
	if err := c.ShouldBindJSON(&deptDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	dept := &sysModel.SysDept{}
	deptDto.Convert(dept)

	dept.DelFlag = enmu.EnmuGroupApp.DelFlagDelete.GetCode()
	dept.Status = enmu.EnmuGroupApp.StatusNormal.GetCode()

	if err := deptService.CreateDept(dept); err != nil {
		logger.Error("create user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (d *SysDeptApi) GetDeptById(c *gin.Context) {

}

func (d *SysDeptApi) ListDept(c *gin.Context) {

}

func (d *SysDeptApi) UpdateDept(c *gin.Context) {

}

func (d *SysDeptApi) DeleteDept(c *gin.Context) {

}
