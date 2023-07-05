package system

import (
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
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

func (m *SysConfigService) GetConfig(params *sysDto.SysConfigQuery) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysConfig{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	var configs []sysModel.SysConfig
	res := db.Limit(limit).Offset(offset).Find(&configs)

	pageResult.Total = total
	pageResult.Rows = configs

	return pageResult, res.Error
}
