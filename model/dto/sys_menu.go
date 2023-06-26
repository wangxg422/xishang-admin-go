package dto

import (
	"backend/model/system"
	"time"
)

type SysCreateMenuDTO struct {
	MenuId     int64     `json:"menuId,omitempty"`
	MenuName   string    `json:"menuName,omitempty" binding:"required"`
	ParentId   int64     `json:"parentId,omitempty"`
	OrderNum   int8      `json:"order_num,omitempty"`
	Path       string    `json:"leader,omitempty" binding:"required"`
	Component  string    `json:"component,omitempty"`
	Query      string    `json:"query,omitempty"`
	IsFrame    string    `json:"isFrame,omitempty"`
	IsCache    string    `json:"isCache,omitempty"`
	MenuType   string    `json:"menuType,omitempty" binding:"required"`
	Visible    string    `json:"visible,omitempty" binding:"required"`
	Status     string    `json:"status,omitempty"`
	Perms      string    `json:"perms,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysCreateMenuDTO) Convert(t *system.SysMenu) {
	t.MenuId = m.MenuId
	t.MenuName = m.MenuName
	t.ParentId = m.ParentId
	t.OrderNum = m.OrderNum
	t.Path = m.Path
	t.Component = m.Component
	t.Query = m.Query
	t.IsFrame = m.IsFrame
	t.IsCache = m.IsCache
	t.MenuType = m.MenuType
	t.Visible = m.Visible
	t.Status = m.Status
	t.Perms = m.Perms
	t.Icon = m.Icon
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysUpdateMenuDTO struct {
	MenuId     int64     `json:"menuId,omitempty"`
	MenuName   string    `json:"menuName,omitempty" binding:"required"`
	ParentId   int64     `json:"parentId,omitempty"`
	OrderNum   int8      `json:"order_num,omitempty"`
	Path       string    `json:"leader,omitempty" binding:"required"`
	Component  string    `json:"component,omitempty"`
	Query      string    `json:"query,omitempty"`
	IsFrame    string    `json:"isFrame,omitempty"`
	IsCache    string    `json:"isCache,omitempty"`
	MenuType   string    `json:"menuType,omitempty" binding:"required"`
	Visible    string    `json:"visible,omitempty" binding:"required"`
	Status     string    `json:"status,omitempty"`
	Perms      string    `json:"perms,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysUpdateMenuDTO) Convert(t *system.SysMenu) {
	t.MenuId = m.MenuId
	t.MenuName = m.MenuName
	t.ParentId = m.ParentId
	t.OrderNum = m.OrderNum
	t.Path = m.Path
	t.Component = m.Component
	t.Query = m.Query
	t.IsFrame = m.IsFrame
	t.IsCache = m.IsCache
	t.MenuType = m.MenuType
	t.Visible = m.Visible
	t.Status = m.Status
	t.Perms = m.Perms
	t.Icon = m.Icon
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}
