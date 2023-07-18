package system

import (
	"backend/model/system"
	"time"
)

type SysCreateMenuDTO struct {
	MenuId     int64     `json:"menuId,omitempty"`
	ParentId   int64     `json:"parentId,omitempty" binding:"required"`
	Name       string    `json:"name,omitempty" binding:"required"`
	Ancestors  string    `json:"ancestors,omitempty"`
	Sort       int8      `json:"sort,omitempty" binding:"required"`
	Path       string    `json:"leader,omitempty"`
	Component  string    `json:"component,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	Query      string    `json:"query,omitempty"`
	Frame      string    `json:"frame,omitempty"`
	Cached     string    `json:"cached,omitempty"`
	Type       string    `json:"type,omitempty"`
	Visible    string    `json:"visible,omitempty"`
	Status     string    `json:"status,omitempty"`
	Perms      string    `json:"perms,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysCreateMenuDTO) Convert(t *system.SysMenu) {
	t.MenuId = m.MenuId
	t.ParentId = m.ParentId
	t.Name = m.Name
	t.Ancestors = m.Ancestors
	t.Path = m.Path
	t.Component = m.Component
	t.Query = m.Query
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
	ParentId   int64     `json:"parentId,omitempty" binding:"required"`
	Name       string    `json:"name,omitempty" binding:"required"`
	Ancestors  string    `json:"ancestors,omitempty"`
	Sort       int8      `json:"sort,omitempty" binding:"required"`
	Path       string    `json:"leader,omitempty"`
	Component  string    `json:"component,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	Query      string    `json:"query,omitempty"`
	Frame      string    `json:"frame,omitempty"`
	Cached     string    `json:"cached,omitempty"`
	Type       string    `json:"type,omitempty"`
	Visible    string    `json:"visible,omitempty"`
	Status     string    `json:"status,omitempty"`
	Perms      string    `json:"perms,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysUpdateMenuDTO) Convert(t *system.SysMenu) {
	t.MenuId = m.MenuId
	t.ParentId = m.ParentId
	t.Path = m.Path
	t.Component = m.Component
	t.Query = m.Query
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
