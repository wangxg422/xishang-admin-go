package system

import "time"

type SysRole struct {
	RoleId            int64     `gorm:"primaryKey;column:role_id" json:"roleId,omitempty"`
	RoleName          string    `gorm:"column:role_name" json:"roleName,omitempty"`
	RoleKey           string    `gorm:"column:role_key" json:"roleKey,omitempty"`
	RoleSort          int64     `gorm:"column:role_sort" json:"roleSort,omitempty"`
	DataScope         int8      `gorm:"column:data_scope" json:"dataScope,omitempty"`
	MenuCheckStrictly int8      `gorm:"column:menu_check_strictly" json:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly int8      `gorm:"column:dept_check_strictly" json:"deptCheckStrictly,omitempty"`
	Status            string    `gorm:"column:status;default:0" json:"status,omitempty"`
	DelFlag           int8      `gorm:"column:del_flag;default:0" json:"delFlag,omitempty"`
	CreateTime        time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime        time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy          string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy          string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark            string    `gorm:"column:remark" json:"remark,omitempty"`

	SysUsers []SysUser `gorm:"many2many:sys_user_role;foreignKey:RoleId;joinForeignKey:RoleId;references:UserId;joinReferences:UserId;" json:"users,omitempty"`
	SysMenus []SysMenu `gorm:"many2many:sys_menu_role;foreignKey:RoleId;joinForeignKey:RoleId;references:MenuId;joinReferences:MenuId;" json:"menus,omitempty"`
}
