package system

import "time"

type SysDept struct {
	DeptId     int64     `gorm:"column:dept_id" json:"deptId,omitempty"`
	DeptName   string    `gorm:"column:dept_name" json:"deptName,omitempty"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,omitempty"`
	Ancestors  int64     `gorm:"column:ancestors" json:"ancestors,omitempty"`
	OrderNum   int8      `gorm:"column:order_num" json:"order_num,omitempty"`
	Leader     string    `gorm:"column:leader" json:"leader,omitempty"`
	Phone      string    `gorm:"column:phone" json:"phone,omitempty"`
	Email      string    `gorm:"column:email" json:"email,omitempty"`
	Status     string    `gorm:"column:status;default:0" json:"status,omitempty"`
	DelFlag    string    `gorm:"column:del_flag;default:0" json:"delFlag,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
}
