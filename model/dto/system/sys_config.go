package system

import (
	"backend/model/common/request"
	"backend/model/system"
	"time"
)

type SysConfigCreateDTO struct {
	ConfigId    int64     `json:"configId,omitempty"`
	ConfigName  string    `json:"configName,omitempty"`
	ConfigKey   string    `json:"configKey,omitempty"`
	ConfigValue string    `json:"configValue,omitempty"`
	ConfigType  int8      `json:"configType,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty"`
	CreateBy    string    `json:"createBy,omitempty"`
	UpdateBy    string    `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
}

func (m *SysConfigCreateDTO) SysDictTypeUpdateDTO(t *system.SysConfig) {
	t.ConfigId = m.ConfigId
	t.ConfigName = m.ConfigName
	t.ConfigKey = m.ConfigKey
	t.ConfigValue = m.ConfigValue
	t.ConfigType = m.ConfigType
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysConfigUpdateDTO struct {
	ConfigId    int64     `json:"configId,omitempty"`
	ConfigName  string    `json:"configName,omitempty"`
	ConfigKey   string    `json:"configKey,omitempty"`
	ConfigValue string    `json:"configValue,omitempty"`
	ConfigType  int8      `json:"configType,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty"`
	CreateBy    string    `json:"createBy,omitempty"`
	UpdateBy    string    `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
}

func (m *SysConfigUpdateDTO) SysDictTypeUpdateDTO(t *system.SysConfig) {
	t.ConfigId = m.ConfigId
	t.ConfigName = m.ConfigName
	t.ConfigKey = m.ConfigKey
	t.ConfigValue = m.ConfigValue
	t.ConfigType = m.ConfigType
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysConfigQuery struct {
	PageInfo   request.PageInfo
	ConfigName string `form:"configName" json:"configName"`
	ConfigKey  string `form:"configKey" json:"configKey"`
	ConfigType string `form:"configType" json:"configType"`
	BeginTime  string `form:"beginTime" json:"beginTime"`
	EndTime    string `form:"endTime" json:"endTime"`
}
