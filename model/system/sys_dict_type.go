package system

import "time"

type SysDictType struct {
	DictId     int64     `gorm:"primaryKey;column:dict_id" json:"dictId,omitempty"`
	DictName   string    `gorm:"column:dict_name" json:"dictName,omitempty"`
	DictType   string    `gorm:"column:dict_type" json:"dictType,omitempty"`
	Status     int8      `gorm:"column:status;default:0" json:"status,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark" json:"remark,omitempty"`
}
