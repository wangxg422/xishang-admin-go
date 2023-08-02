package system

import (
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysPermissionApi struct {
}

func (m *SysPermissionApi) DeletePermission(c *gin.Context) {
	id := c.Param("permId")
	if id == "" {
		response.FailWithMessage("permId is null", c)
		return
	}

	permId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = permissionService.DeletePermission(permId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPermissionApi) CreatePermission(c *gin.Context) {
	dto := sysDto.SysPermissionCreateDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	perm := &sysModel.SysPermission{}
	dto.Convert(perm)
	perm.CreateBy = jwt.GetUserName(c)

	if err := permissionService.CreatePermission(perm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysPermissionApi) GetPermissionById(c *gin.Context) {
	id := c.Param("permId")
	if id == "" {
		response.FailWithMessage("configId is null", c)
		return
	}

	permId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	perm, err := permissionService.GetPermissionById(permId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(perm, c)
}

func (m *SysPermissionApi) UpdatePermission(c *gin.Context) {
	dto := sysDto.SysPermissionUpdateDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	perm := &sysModel.SysPermission{}
	err := dto.Convert(perm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	perm.UpdateBy = jwt.GetUserName(c)

	if err := permissionService.UpdatePermission(perm); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
