package system

import (
	"backend/common/enmu"
	"backend/common/response"
	"backend/initial/logger"
	"backend/model/dto"
	sysModel "backend/model/system"
	"backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDeptApi struct{}

func (m *SysDeptApi) CreateDept(c *gin.Context) {
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
		logger.Error("create dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDeptApi) GetDeptById(c *gin.Context) {
	id := c.Param("deptId")

	if id == "" {
		response.FailWithMessage("dept id is null", c)
		return
	}

	deptId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("dept id convert failed", c)
		return
	}

	user, err := deptService.GetDeptById(deptId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysDeptApi) ListDept(c *gin.Context) {

}

func (m *SysDeptApi) UpdateDept(c *gin.Context) {
	deptDto := dto.SysUpdateDeptDTO{}

	if err := c.ShouldBindJSON(&deptDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if deptDto.DeptId == 0 {
		response.FailWithMessage("dept id can not be null", c)
		return
	}

	dept := &sysModel.SysDept{}
	deptDto.Convert(dept)
	if err := deptService.UpdateDept(dept); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDeptApi) DeleteDept(c *gin.Context) {
	id := c.Param("deptId")

	deptId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("dept id convert failed", c)
		return
	}

	if err := deptService.DeleteDept(deptId); err != nil {
		logger.Error("delete dept failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
