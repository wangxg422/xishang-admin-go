package system

import "time"

type SysModule struct {
	ModuleId   int64     `gorm:"primaryKey;column:module_id" json:"moduleId"`
	ModuleName string    `gorm:"column:module_name" json:"moduleName"`
	ModuleCode string    `gorm:"column:module_Code" json:"moduleCode"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`
	Remark     string    `gorm:"column:remark" json:"remark"`
}
