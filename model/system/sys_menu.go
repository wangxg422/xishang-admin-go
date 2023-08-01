package system

import "time"

type SysMenu struct {
	MenuId     int64     `gorm:"primaryKey;column:menu_id" json:"menuId,string"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,string"`
	MenuCode   string    `gorm:"column:menu_code" json:"menuCode"`
	MenuName   string    `gorm:"column:menu_name" json:"menuName"`
	Ancestors  string    `gorm:"column:ancestors" json:"ancestors"`
	Sort       int8      `gorm:"column:sort" json:"sort"`
	Path       string    `gorm:"column:path" json:"path"`
	Component  string    `gorm:"column:component" json:"component"`
	Icon       string    `gorm:"column:icon" json:"icon"`
	Query      string    `gorm:"column:query" json:"query"`
	Frame      string    `gorm:"column:frame" json:"frame"`   // 是否为外链
	Cached     string    `gorm:"column:cached" json:"cached"` // 是否缓存
	Type       string    `gorm:"column:type" json:"type"`     // M目录 C菜单 F按钮
	Hidden     string    `gorm:"column:hidden" json:"hidden"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy"`
	Remark     string    `gorm:"column:remark" json:"remark"`

	SysRoles []SysRole `gorm:"many2many:sys_role_menu;foreignKey:MenuId;joinForeignKey:MenuId;references:RoleId;joinReferences:RoleId;" json:"roles,omitempty"`
	// 子菜单
	Children []SysMenu `gorm:"-" json:"children,omitempty"`
}
