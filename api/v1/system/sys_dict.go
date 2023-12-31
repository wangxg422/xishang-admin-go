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
		logger.Error("查询字典类型失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(types, c)
}

func (m *SysDictApi) CreateDictType(c *gin.Context) {
	params := &sysDto.SysDictTypeCreateDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	dictType := &sysModel.SysDictType{}
	params.Convert(dictType)

	dictType.CreateBy = jwt.GetUserName(c)

	err = dictService.CreateDictType(dictType)
	if err != nil {
		logger.Error("创建字典类型失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDictApi) UpdateDictType(c *gin.Context) {
	params := &sysDto.SysDictTypeUpdateDTO{}
	err := c.ShouldBind(params)

	if params.DictTypeId == 0 {
		response.FailWithMessage("dictTypeId is null", c)
		return
	}

	dictType := &sysModel.SysDictType{}
	params.Convert(dictType)
	dictType.UpdateBy = jwt.GetUserName(c)
	dictType.UpdateTime = time.Now()

	err = dictService.UpdateDictType(dictType)
	if err != nil {
		logger.Error("update dictType failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDictApi) GetDictTypeById(c *gin.Context) {
	typeIdStr := c.Param("dictTypeId")
	if typeIdStr == "" {
		response.FailWithMessage("dictTypeId is null", c)
		return
	}

	dictTypeId, err := strconv.ParseInt(typeIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	dictType, err := dictService.GetDictTypeById(dictTypeId)
	if err != nil {
		logger.Error("get dict type failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(dictType, c)
}

func (m *SysDictApi) DeleteDictType(c *gin.Context) {
	dictTypeIdStr := c.Param("dictTypeId")
	if dictTypeIdStr == "" {
		response.FailWithMessage("dictTypeId is null", c)
		return
	}

	ids := strings.Split(dictTypeIdStr, constant.Comma)
	dictTypeIds, err := utils.StrToInt64Array(ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = dictService.DeleteDictType(dictTypeIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDictApi) GetDictTypeAll(c *gin.Context) {
	dictTypes, err := dictService.GetDictTypeAll()
	if err != nil {
		logger.Error("get dict type failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(dictTypes, c)
}

func (m *SysDictApi) GetDictDataPage(c *gin.Context) {
	params := &sysDto.SysDictDataQueryDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	datas, err := dictService.GetDictDataPage(params)
	if err != nil {
		logger.Error("查询字典数据失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(datas, c)
}

func (m *SysDictApi) CreateDictData(c *gin.Context) {
	params := &sysDto.SysDictDataCreateDTO{}
	err := c.ShouldBind(params)

	if err != nil {
		logger.Error("参数解析失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	dictData := &sysModel.SysDictData{}
	params.Convert(dictData)

	dictData.CreateBy = jwt.GetUserName(c)

	err = dictService.CreateDictData(dictData)
	if err != nil {
		logger.Error("创建字典数据失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDictApi) UpdateDictData(c *gin.Context) {
	params := &sysDto.SysDictDataUpdateDTO{}
	err := c.ShouldBind(params)

	if params.DictDataId == 0 {
		response.FailWithMessage("dictDataId is null", c)
		return
	}

	dictData := &sysModel.SysDictData{}
	params.Convert(dictData)
	dictData.UpdateBy = jwt.GetUserName(c)

	err = dictService.UpdateDictData(dictData)
	if err != nil {
		logger.Error("update dictData failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}

func (m *SysDictApi) GetDictDataById(c *gin.Context) {
	dataIdStr := c.Param("dictDataId")
	if dataIdStr == "" {
		response.FailWithMessage("dictDataId is null", c)
		return
	}

	dictDataId, err := strconv.ParseInt(dataIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	dictData, err := dictService.GetDictDataById(dictDataId)
	if err != nil {
		logger.Error("get dict data failed", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(dictData, c)
}

func (m *SysDictApi) DeleteDictData(c *gin.Context) {
	dictDataIdStr := c.Param("dictDataId")
	if dictDataIdStr == "" {
		response.FailWithMessage("dictDataId is null", c)
		return
	}

	ids := strings.Split(dictDataIdStr, constant.Comma)
	dictDataIds, err := utils.StrToInt64Array(ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = dictService.DeleteDictData(dictDataIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
}
