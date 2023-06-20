package system

import (
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"
	sysSvc "backend/service/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDeptApi struct{}

var deptSvc = sysSvc.SysDeptService{}
var userSvc = sysSvc.SysUserService{}

func (d *SysDeptApi) CreateDept(c *gin.Context) {
	deptDto := dto.SysCreateDeptDTO{}
	if err := c.ShouldBindJSON(&deptDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	dept := &sysModel.SysDept{}
	deptDto.Convert(dept)

	if err := deptSvc.CreateDept(dept); err != nil {
		logger.Error("create user failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
