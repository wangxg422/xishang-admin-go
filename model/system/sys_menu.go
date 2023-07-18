package system

import "time"

type SysMenu struct {
	MenuId     int64     `gorm:"primaryKey;column:menu_id" json:"menuId,omitempty"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,omitempty"`
	Name       string    `gorm:"column:name" json:"name,omitempty"`
	Title      string    `gorm:"column:title" json:"title,omitempty"`
	Ancestors  string    `gorm:"column:ancestors" json:"ancestors,omitempty"`
	Sort       int8      `gorm:"column:sort" json:"sort,omitempty"`
	Path       string    `gorm:"column:path" json:"leader,omitempty"`
	Component  string    `gorm:"column:component" json:"component,omitempty"`
	Icon       string    `gorm:"column:icon" json:"icon,omitempty"`
	Query      string    `gorm:"column:query" json:"query,omitempty"`
	Frame      string    `gorm:"column:frame" json:"frame,omitempty"`   // 是否为外链
	Cached     string    `gorm:"column:cached" json:"cached,omitempty"` // 是否缓存
	Type       string    `gorm:"column:type" json:"type,omitempty"`     // M目录 C菜单 F按钮
	Hidden     string    `gorm:"column:hidden" json:"hidden,omitempty"`
	Status     string    `gorm:"column:status;default:0" json:"status,omitempty"`
	Perms      string    `gorm:"column:perms" json:"perms,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark;default:0" json:"remark,omitempty"`

	SysRoles []SysRole `gorm:"many2many:sys_menu_role;foreignKey:MenuId;joinForeignKey:menu_id;references:RoleId;joinReferences:role_id;" json:"roles,omitempty"`
	// 子菜单
	Children []SysMenu `gorm:"-" json:"children,omitempty"`
}
