package system

import (
	"backend/common/enmu"
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	"backend/utils"
	"backend/utils/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysRoleApi struct{}

func (m *SysRoleApi) CreateRole(c *gin.Context) {
	roleDto := sysDto.SysRoleCreateDTO{}
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	role := roleDto.Convert()

	role.DelFlag = enmu.DelFlagDeleted.Value()
	role.Status = enmu.StatusNormal.Value()
	role.CreateBy = jwt.GetUserName(c)

	if err := roleService.CreateRole(&role); err != nil {
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

func (m *SysRoleApi) GetRolePage(c *gin.Context) {
	params := &sysDto.SysRoleQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	roles, err := roleService.GetRolePage(params)
	if err != nil {
		logger.Error("查询角色失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(roles, c)
}

func (m *SysRoleApi) UpdateRole(c *gin.Context) {
	roleDto := sysDto.SysRoleUpdateDTO{}

	if err := c.ShouldBindJSON(&roleDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if roleDto.RoleId == "" {
		response.FailWithMessage("role id can not be null", c)
		return
	}

	role, err := roleDto.Convert()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.UpdateRole(&role); err != nil {
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
