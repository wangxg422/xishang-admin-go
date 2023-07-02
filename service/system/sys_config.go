package system

import (
	"backend/global"
	"backend/model/system"
)

type SysConfigService struct {
}

func (m *SysConfigService) GetConfigByKey(key string) (system.SysConfig, error) {
	var config system.SysConfig
	res := global.DB.
		Where("config_key = ?", key).Order("dict_sort").
		Find(&config)

	return config, res.Error
}
