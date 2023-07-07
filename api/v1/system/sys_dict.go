package system

import (
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysDictApi struct {
}

func (m *SysDictApi) GetDictDataByType(c *gin.Context) {
	dictType := c.Param("dictType")

	if dictType == "" {
		response.FailWithMessage("dict type is null", c)
		return
	}

	data, err := dictService.GetDictDataByType(dictType)
	if err != nil {
		logger.Error("search dict failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(data, c)
}

func (m *SysDictApi) GetDictTypePage(c *gin.Context) {
	params := &sysDto.SysDictTypeQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	types, err := dictService.GetDictTypePage(params)
	if err != nil {
		logger.Error("查询配置失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(types, c)
}
