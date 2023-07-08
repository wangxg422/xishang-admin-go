package system

import (
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
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

	likeArr := []string{
		"config_name",
		"config_key",
	}

	utils.ConcatLikeWhereCondition(db, likeArr, params.ConfigName, params.ConfigKey)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)
	utils.ConcatOneEqualsStrWhereCondition(db, "inner_config", params.InnerConfig)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	var configs []sysModel.SysConfig
	res := db.Find(&configs)

	pageResult.Total = total
	pageResult.Rows = configs

	return pageResult, res.Error
}

func (m *SysConfigService) CreateConfig(config *sysModel.SysConfig) error {
	var existConfig sysModel.SysConfig
	err := global.DB.
		Where("config_key = ?", config.ConfigKey).
		Find(&existConfig).Limit(1).Error
	if err != nil {
		return err
	}

	if existConfig.ConfigId != 0 {
		return errors.New("系统配置 " + config.ConfigKey + " 已经存在")
	}

	return global.DB.Create(config).Error
}

func (m *SysConfigService) GetConfigById(configId int64) (sysModel.SysConfig, error) {
	var config sysModel.SysConfig
	res := global.DB.
		Where("config_id = ?", configId).
		Find(&config)

	return config, res.Error
}

func (m *SysConfigService) GetConfigByIds(configIds []int64) ([]sysModel.SysConfig, error) {
	var config []sysModel.SysConfig
	res := global.DB.
		Where("config_Id IN ?", configIds).
		Find(&config)

	return config, res.Error
}

func (m *SysConfigService) UpdateConfig(config *sysModel.SysConfig) error {
	// 这里保证零值也能更新
	return global.DB.Select("config_name", "config_key", "config_value",
		"inner_config", "update_time", "update_by", "remark").
		Updates(config).Error
}

func (m *SysConfigService) DeleteConfig(configIds []int64) error {
	// 系统内置配置禁止删除
	configs, err := m.GetConfigByIds(configIds)
	if err != nil {
		return err
	}

	for _, c := range configs {
		if c.InnerConfig == "1" {
			return errors.New("系统内置配置禁止删除")
		}
	}

	return global.DB.Where("config_id IN ?", configIds).Delete(&sysModel.SysConfig{}).Error
}
