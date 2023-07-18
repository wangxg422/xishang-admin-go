package system

import "time"

type SysMenu struct {
	MenuId     int64     `gorm:"primaryKey;column:menu_id" json:"menuId,string"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,string"`
	Name       string    `gorm:"column:name" json:"name"`
	Title      string    `gorm:"column:title" json:"title"`
	Ancestors  string    `gorm:"column:ancestors" json:"ancestors"`
	Sort       int8      `gorm:"column:sort" json:"sort"`
	Path       string    `gorm:"column:path" json:"leader"`
	Component  string    `gorm:"column:component" json:"component"`
	Icon       string    `gorm:"column:icon" json:"icon"`
	Query      string    `gorm:"column:query" json:"query"`
	Frame      string    `gorm:"column:frame" json:"frame"`   // 是否为外链
	Cached     string    `gorm:"column:cached" json:"cached"` // 是否缓存
	Type       string    `gorm:"column:type" json:"type"`     // M目录 C菜单 F按钮
	Hidden     string    `gorm:"column:hidden" json:"hidden"`
	Status     string    `gorm:"column:status;default:0" json:"status"`
	Perms      string    `gorm:"column:perms" json:"perms"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`
	Remark     string    `gorm:"column:remark;default:0" json:"remark"`

	SysRoles []SysRole `gorm:"many2many:sys_menu_role;foreignKey:MenuId;joinForeignKey:menu_id;references:RoleId;joinReferences:role_id;" json:"roles,omitempty"`
	// 子菜单
	Children []SysMenu `gorm:"-" json:"children,omitempty"`
}
