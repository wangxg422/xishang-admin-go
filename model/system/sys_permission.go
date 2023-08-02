package system

import "time"

type SysPermission struct {
	PermId     int64     `gorm:"primaryKey;column:perm_id" json:"permId,string"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,string"`
	ModuleCode string    `gorm:"column:module_code" json:"moduleCode"`
	PermName   string    `gorm:"column:perm_name" json:"permName"`
	PermCode   string    `gorm:"column:post_code" json:"postCode"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`
	Remark     string    `gorm:"column:remark" json:"remark"`

	SysRoles []SysRole `gorm:"many2many:sys_role_permission;foreignKey:PermId;joinForeignKey:PermId;references:RoleId;joinReferences:RoleId;" json:"roles,omitempty"`
}
