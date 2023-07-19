package system

import "time"

type SysRole struct {
	RoleId            int64     `gorm:"primaryKey;column:role_id" json:"roleId,string"`
	RoleName          string    `gorm:"column:role_name" json:"roleName"`
	RoleCode          string    `gorm:"column:role_code" json:"roleCode"`
	RolePerms         string    `gorm:"column:role_perm" json:"rolePerms"`
	Sort              int64     `gorm:"column:sort" json:"sort"`
	DataScope         int8      `gorm:"column:data_scope" json:"dataScope"`
	MenuCheckStrictly int8      `gorm:"column:menu_check_strictly" json:"menuCheckStrictly"`
	DeptCheckStrictly int8      `gorm:"column:dept_check_strictly" json:"deptCheckStrictly"`
	Status            string    `gorm:"column:status" json:"status"`
	DelFlag           int8      `gorm:"column:del_flag" json:"delFlag"`
	CreateTime        time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime        time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy          string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy          string    `gorm:"column:update_by" json:"updateBy"`
	Remark            string    `gorm:"column:remark" json:"remark"`

	SysUsers []SysUser `gorm:"many2many:sys_user_role;foreignKey:RoleId;joinForeignKey:RoleId;references:UserId;joinReferences:UserId;" json:"users,omitempty"`
	SysMenus []SysMenu `gorm:"many2many:sys_menu_role;foreignKey:RoleId;joinForeignKey:RoleId;references:MenuId;joinReferences:MenuId;" json:"menus,omitempty"`
	SysDepts []SysDept `gorm:"many2many:sys_role_dept;foreignKey:RoleId;joinForeignKey:RoleId;references:DeptId;joinReferences:DeptId;" json:"depts,omitempty"`
}
