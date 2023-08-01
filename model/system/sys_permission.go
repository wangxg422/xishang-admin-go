package system

import "time"

type SysPermission struct {
	PermId     int64     `gorm:"primaryKey;column:perm_id" json:"permId"`
	ModuleId   string    `gorm:"column:module_id" json:"moduleId"`
	PermName   string    `gorm:"column:perm_name" json:"permName"`
	PermCode   string    `gorm:"column:post_code" json:"postCode"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`
	Remark     string    `gorm:"column:remark" json:"remark"`
}
