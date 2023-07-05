package system

import (
	"backend/global"
	"backend/model/common/request"
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

func (m *SysConfigService) GetConfig(page request.PageInfo) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysConfig{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	var configs []sysModel.SysConfig
	res := db.Limit(page.Limit).Offset(page.Offset).Find(&configs)

	pageResult.Total = total
	pageResult.List = configs

	return pageResult, res.Error
}
