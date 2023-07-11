package system

import (
	"backend/common/enmu"
	"backend/initial/logger"
	"backend/model/common/response"
	"backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysRoleApi struct{}

func (m *SysRoleApi) CreateRole(c *gin.Context) {
	roleDto := system.SysCreateRoleDTO{}
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	role := &sysModel.SysRole{}
	roleDto.Convert(role)

	role.DelFlag = enmu.DelFlagDeleted.Value()
	role.Status = enmu.StatusNormal.Value()

	if err := roleService.CreateRole(role); err != nil {
		logger.Error("create role failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysRoleApi) GetRoleById(c *gin.Context) {
	id := c.Param("roleId")

	if id == "" {
		response.FailWithMessage("role id is null", c)
		return
	}

	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("role id convert failed", c)
		return
	}

	user, err := roleService.GetRoleById(roleId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search role failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysRoleApi) ListRole(c *gin.Context) {

}

func (m *SysRoleApi) UpdateRole(c *gin.Context) {
	roleDto := system.SysUpdateRoleDTO{}

	if err := c.ShouldBindJSON(&roleDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if roleDto.RoleId == 0 {
		response.FailWithMessage("role id can not be null", c)
		return
	}

	role := &sysModel.SysRole{}
	roleDto.Convert(role)
	if err := roleService.UpdateRole(role); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysRoleApi) DeleteRole(c *gin.Context) {
	id := c.Param("roleId")

	roleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("role id convert failed", c)
		return
	}

	if err := roleService.DeleteRole(roleId); err != nil {
		logger.Error("delete role failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysRoleApi) GetAllRole(c *gin.Context) {
	roles, err := roleService.GetAllRole()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(roles, c)
}
