package system

import "time"

type SysMenu struct {
	MenuId     int64     `gorm:"primaryKey;column:menu_id" json:"menuId,omitempty"`
	MenuName   string    `gorm:"column:menu_name" json:"menuName,omitempty"`
	ParentId   int64     `gorm:"column:parent_id" json:"parentId,omitempty"`
	OrderNum   int8      `gorm:"column:order_num" json:"order_num,omitempty"`
	Path       string    `gorm:"column:path" json:"leader,omitempty"`
	Component  string    `gorm:"column:component" json:"component,omitempty"`
	Query      string    `gorm:"column:query" json:"query,omitempty"`
	IsFrame    int8      `gorm:"column:is_frame" json:"isFrame,omitempty"`   // 是否为外链
	IsCache    int8      `gorm:"column:is_cache" json:"isCache,omitempty"`   // 是否缓存
	MenuType   string    `gorm:"column:menu_type" json:"menuType,omitempty"` // M目录 C菜单 F按钮
	Visible    int8      `gorm:"column:visible" json:"visible,omitempty"`
	Status     int8      `gorm:"column:status;default:0" json:"status,omitempty"`
	Perms      string    `gorm:"column:perms" json:"perms,omitempty"`
	Icon       string    `gorm:"column:icon" json:"icon,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark;default:0" json:"remark,omitempty"`

	SysRoles []SysRole `gorm:"many2many:sys_menu_role;foreignKey:MenuId;joinForeignKey:menu_id;references:RoleId;joinReferences:role_id;" json:"roles,omitempty"`
	// 子菜单
	Children []SysMenu `json:"children,omitempty"`
}
