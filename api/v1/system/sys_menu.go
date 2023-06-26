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

type SysMenuApi struct{}

func (m *SysMenuApi) CreateMenu(c *gin.Context) {
	menuDto := dto.SysCreateMenuDTO{}
	if err := c.ShouldBindJSON(&menuDto); err != nil {
		logger.Error("param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	menu := &sysModel.SysMenu{}
	menuDto.Convert(menu)

	menu.Status = enmu.EnmuGroupApp.StatusNormal.GetCode()

	if err := menuService.CreateMenu(menu); err != nil {
		logger.Error("create menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) GetMenuById(c *gin.Context) {
	id := c.Param("menuId")

	if id == "" {
		response.FailWithMessage("menu id is null", c)
		return
	}

	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("menu id convert failed", c)
		return
	}

	user, err := menuService.GetMenuById(menuId)
	if err != nil {
		if utils.NoRecord(err) {
			response.OkWithData([]string{}, c)
			return
		}
		logger.Error("search menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(user, c)
}

func (m *SysMenuApi) ListMenu(c *gin.Context) {

}

func (m *SysMenuApi) UpdateMenu(c *gin.Context) {
	menuDto := dto.SysUpdateMenuDTO{}

	if err := c.ShouldBindJSON(&menuDto); err != nil {
		logger.Error("parse param error", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if menuDto.MenuId == 0 {
		response.FailWithMessage("menu id can not be null", c)
		return
	}

	menu := &sysModel.SysMenu{}
	menuDto.Convert(menu)
	if err := menuService.UpdateMenu(menu); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysMenuApi) DeleteMenu(c *gin.Context) {
	id := c.Param("menuId")

	menuId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailWithMessage("menu id convert failed", c)
		return
	}

	if err := menuService.DeleteMenu(menuId); err != nil {
		logger.Error("delete menu failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
