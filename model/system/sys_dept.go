package system

import "time"

type SysDept struct {
	DeptId      int64     `gorm:"primaryKey;column:dept_id" json:"deptId,string"`
	DeptName    string    `gorm:"column:dept_name" json:"deptName"`
	DeptCode    string    `gorm:"column:dept_code" json:"deptCode"`
	ParentId    int64     `gorm:"column:parent_id" json:"parentId,string"`
	Ancestors   string    `gorm:"column:ancestors" json:"ancestors"`
	Sort        int8      `gorm:"column:sort" json:"sort"`
	Leader      string    `gorm:"column:leader" json:"leader"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	Email       string    `gorm:"column:email" json:"email"`
	Status      string    `gorm:"column:status" json:"status"`
	DelFlag     int8      `gorm:"column:del_flag" json:"delFlag"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy    string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy    string    `gorm:"column:update_by" json:"updateBy"`

	SysRoles []SysRole `gorm:"many2many:sys_role_dept;foreignKey:DeptId;joinForeignKey:DeptId;references:RoleId;joinReferences:RoleId;" json:"roles,omitempty"`
}
