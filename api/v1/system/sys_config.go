package system

import (
	"backend/common/constant"
	"backend/initial/logger"
	"backend/model/common/response"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"backend/utils/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
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

func (m *SysConfigApi) GetConfigPage(c *gin.Context) {
	params := &sysDto.SysConfigQuery{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	configs, err := configService.GetConfigPage(params)
	if err != nil {
		logger.Error("查询配置失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(configs, c)
}

func (m *SysConfigApi) CreateConfig(c *gin.Context) {
	params := &sysDto.SysConfigCreateDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	config := &sysModel.SysConfig{}
	params.Convert(config)

	config.CreateBy = jwt.GetUserName(c)

	err = configService.CreateConfig(config)
	if err != nil {
		logger.Error("创建配置失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysConfigApi) GetConfigById(c *gin.Context) {
	configIdStr := c.Param("configId")
	if configIdStr == "" {
		response.FailWithMessage("configId is null", c)
		return
	}

	configId, err := strconv.ParseInt(configIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	config, err := configService.GetConfigById(configId)
	if err != nil {
		logger.Error("get config failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(config, c)
}

func (m *SysConfigApi) UpdateConfig(c *gin.Context) {
	params := &sysDto.SysConfigUpdateDTO{}
	err := c.ShouldBind(params)

	if params.ConfigId == 0 {
		response.FailWithMessage("configId is null", c)
		return
	}

	config := &sysModel.SysConfig{}
	params.Convert(config)
	config.UpdateBy = jwt.GetUserName(c)
	config.UpdateTime = time.Now()

	err = configService.UpdateConfig(config)
	if err != nil {
		logger.Error("update config failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysConfigApi) DeleteConfig(c *gin.Context) {
	configIdStr := c.Param("configId")
	if configIdStr == "" {
		response.FailWithMessage("configId is null", c)
		return
	}

	ids := strings.Split(configIdStr, constant.Comma)
	configIds, err := utils.StrToInt64Array(ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = configService.DeleteConfig(configIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
