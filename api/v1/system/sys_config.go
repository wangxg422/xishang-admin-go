package system

import (
	"backend/initial/logger"
	"backend/model/common/request"
	"backend/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysConfigApi struct {
}

func (m *SysConfigApi) GetConfigByKey(c *gin.Context) {
	configKey := c.Param("configKey")

	if configKey == "" {
		response.FailWithMessage("config key is null", c)
		return
	}

	value, err := configService.GetConfigByKey(configKey)
	if err != nil {
		logger.Error("search config failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if value.ConfigValue == "" {
		logger.Error("config %s not exist", zap.String("key", configKey))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(value.ConfigValue, c)
}

func (m *SysConfigApi) GetConfig(c *gin.Context) {
	pageInfo, err := request.NewPageInfo(c)
	if err != nil {
		logger.Error("获取分页信息失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	configs, err := configService.GetConfig(pageInfo)
	if err != nil {
		logger.Error("查询配置失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(configs, c)
}
