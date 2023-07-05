package system

import (
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
)

type SysConfigService struct {
}

func (m *SysConfigService) GetConfigByKey(key string) (sysModel.SysConfig, error) {
	var config sysModel.SysConfig
	res := global.DB.
		Where("config_key = ?", key).
		Find(&config)

	return config, res.Error
}

func (m *SysConfigService) GetConfigPage(params *sysDto.SysConfigQuery) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysConfig{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	likeArr := []string{
		"config_name",
		"config_key",
	}
	utils.ConcatLikeWhereCondition(db, likeArr, params.ConfigName, params.ConfigKey)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)

	if params.ConfigType != "" {
		db.Where("config_type = ?", params.ConfigType)
	}

	var configs []sysModel.SysConfig
	res := db.Find(&configs)

	pageResult.Total = total
	pageResult.Rows = configs

	return pageResult, res.Error
}
