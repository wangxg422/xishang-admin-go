package system

import "time"

type SysRole struct {
	RoleId            int64     `gorm:"column:role_id" json:"roleId,omitempty"`
	RoleName          string    `gorm:"column:role_name" json:"roleName,omitempty"`
	RoleKey           string    `gorm:"column:role_key" json:"roleKey,omitempty"`
	RoleSort          int64     `gorm:"column:role_sort" json:"roleSort,omitempty"`
	DataScope         int8      `gorm:"column:data_scope" json:"dataScope,omitempty"`
	MenuCheckStrictly string    `gorm:"column:menu_check_strictly" json:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly string    `gorm:"column:dept_check_strictly" json:"deptCheckStrictly,omitempty"`
	Status            string    `gorm:"column:status;default:0" json:"status,omitempty"`
	DelFlag           string    `gorm:"column:del_flag;default:0" json:"delFlag,omitempty"`
	CreateTime        time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime        time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy          string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy          string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark            string    `gorm:"column:remark" json:"remark,omitempty"`
}
