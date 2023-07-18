package system

import (
	"backend/model/system"
	"strconv"
	"time"
)

type SysMenuCreateDTO struct {
	ParentId  string `json:"parentId" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Sort      int8   `json:"sort" binding:"required"`
	Path      string `json:"path" binding:"required"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Query     string `json:"query"`
	Frame     string `json:"frame" binding:"required"`
	Cached    string `json:"cached" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Hidden    string `json:"hidden" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Perms     string `json:"perms"`
	Remark    string `json:"remark"`
}

func (m *SysMenuCreateDTO) Convert(t *system.SysMenu) error {
	id, err := strconv.ParseInt(m.ParentId, 10, 64)
	if err != nil {
		return err
	}

	t.ParentId = id
	t.Name = m.Name
	t.Title = m.Title
	t.Sort = m.Sort
	t.Path = m.Path
	t.Component = m.Component
	t.Icon = m.Icon
	t.Query = m.Query
	t.Frame = m.Frame
	t.Cached = m.Cached
	t.Type = m.Type
	t.Hidden = m.Hidden
	t.Status = m.Status
	t.Perms = m.Perms
	t.Remark = m.Remark
	return nil
}

type SysUpdateMenuDTO struct {
	MenuId     int64     `json:"menuId,omitempty"`
	ParentId   int64     `json:"parentId,omitempty" binding:"required"`
	Name       string    `json:"name,omitempty" binding:"required"`
	Title      string    `json:"title,omitempty" binding:"required"`
	Ancestors  string    `json:"ancestors,omitempty"`
	Sort       int8      `json:"sort,omitempty" binding:"required"`
	Path       string    `json:"leader,omitempty"`
	Component  string    `json:"component,omitempty"`
	Icon       string    `json:"icon,omitempty"`
	Query      string    `json:"query,omitempty"`
	Frame      string    `json:"frame,omitempty"`
	Cached     string    `json:"cached,omitempty"`
	Type       string    `json:"type,omitempty"`
	Hidden     string    `json:"hidden,omitempty"`
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
	t.Hidden = m.Hidden
	t.Status = m.Status
	t.Perms = m.Perms
	t.Icon = m.Icon
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysMenuQuery struct {
	Name   string `form:"name" json:"name"`
	Title  string `form:"title" json:"title"`
	Status string `form:"status" json:"status"`
}
