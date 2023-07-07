package system

import (
	"backend/model/common/request"
	"backend/model/system"
)

type SysConfigCreateDTO struct {
	ConfigName  string `json:"configName,omitempty" binding:"required"`
	ConfigKey   string `json:"configKey,omitempty" binding:"required"`
	ConfigValue string `json:"configValue,omitempty" binding:"required"`
	InnerConfig string `json:"innerConfig,omitempty" binding:"required"`
	Remark      string `json:"remark,omitempty"`
}

func (m *SysConfigCreateDTO) Convert(t *system.SysConfig) {
	t.ConfigName = m.ConfigName
	t.ConfigKey = m.ConfigKey
	t.ConfigValue = m.ConfigValue
	t.InnerConfig = m.InnerConfig
	t.Remark = m.Remark
}

type SysConfigUpdateDTO struct {
	ConfigId    int64  `json:"configId,omitempty"`
	ConfigName  string `json:"configName,omitempty" binding:"required"`
	ConfigKey   string `json:"configKey,omitempty" binding:"required"`
	ConfigValue string `json:"configValue,omitempty" binding:"required"`
	InnerConfig string `json:"innerConfig,omitempty" binding:"required"`
	Remark      string `json:"remark,omitempty"`
}

func (m *SysConfigUpdateDTO) Convert(t *system.SysConfig) {
	t.ConfigId = m.ConfigId
	t.ConfigName = m.ConfigName
	t.ConfigKey = m.ConfigKey
	t.ConfigValue = m.ConfigValue
	t.InnerConfig = m.InnerConfig
	t.Remark = m.Remark
}

type SysConfigQuery struct {
	PageInfo    request.PageInfo
	ConfigName  string `form:"configName" json:"configName"`
	ConfigKey   string `form:"configKey" json:"configKey"`
	InnerConfig string `form:"innerConfig" json:"innerConfig"`
	BeginTime   string `form:"beginTime" json:"beginTime"`
	EndTime     string `form:"endTime" json:"endTime"`
}
