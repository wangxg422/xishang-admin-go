package system

import "time"

type SysDictType struct {
	DictTypeId int64     `gorm:"primaryKey;column:dict_type_id" json:"dictTypeId,omitempty"`
	DictName   string    `gorm:"column:dict_name" json:"dictName,omitempty"`
	DictType   string    `gorm:"column:dict_type" json:"dictType,omitempty"`
	Status     string    `gorm:"column:status;default:0" json:"status,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark" json:"remark,omitempty"`
}
