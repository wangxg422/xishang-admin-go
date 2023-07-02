package system

import "time"

type SysConfig struct {
	ConfigId    int64     `gorm:"primaryKey;column:config_id" json:"configId,omitempty"`
	ConfigName  string    `gorm:"column:config_name" json:"configName,omitempty"`
	ConfigKey   string    `gorm:"column:config_key" json:"configKey,omitempty"`
	ConfigValue string    `gorm:"column:config_value" json:"configValue,omitempty"`
	ConfigType  string    `gorm:"column:config_type;default:N" json:"configType,omitempty"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy    string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy    string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark      string    `gorm:"column:remark" json:"remark,omitempty"`
}
