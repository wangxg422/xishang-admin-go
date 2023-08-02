package system

import (
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type SysModuleApi struct {
}

func (m *SysModuleApi) CreateModule(c *gin.Context) {
	dto := sysDto.SysModuleCreateDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mod := &sysModel.SysModule{}
	dto.Convert(mod)
	mod.CreateBy = jwt.GetUserName(c)

	if err := moduleService.CreateModule(mod); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysModuleApi) GetModule(c *gin.Context) {
	mods, err := moduleService.GetModule()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(mods, c)
}

func (m *SysModuleApi) GetModuleById(c *gin.Context) {
	id := c.Param("moduleId")
	if id == "" {
		response.FailWithMessage("module is null", c)
		return
	}

	modId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mods, err := moduleService.GetModuleById(modId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(mods, c)
}

func (m *SysModuleApi) DeleteModule(c *gin.Context) {
	id := c.Param("moduleId")
	if id == "" {
		response.FailWithMessage("moduleId is null", c)
		return
	}

	modId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = moduleService.DeleteModule(modId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysModuleApi) UpdateModule(c *gin.Context) {
	dto := sysDto.SysModuleUpdateDTO{}
	if err := c.ShouldBindJSON(&dto); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mod := &sysModel.SysModule{}
	err := dto.Convert(mod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mod.UpdateBy = jwt.GetUserName(c)

	if err := moduleService.UpdateModule(mod); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
