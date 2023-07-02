package system

import (
	"backend/common/response"
	"backend/initial/logger"
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
