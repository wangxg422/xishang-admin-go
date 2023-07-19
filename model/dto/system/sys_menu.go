package system

import (
	"backend/model/system"
	"strconv"
)

type SysMenuCreateDTO struct {
	ParentId  string `json:"parentId" binding:"required"`
	MenuName  string `json:"menuName" binding:"required"`
	MenuCode  string `json:"menuCode" binding:"required"`
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
	t.MenuName = m.MenuName
	t.MenuCode = m.MenuCode
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

type SysMenuUpdateDTO struct {
	MenuId    string `json:"menuId" binding:"required"`
	ParentId  string `json:"parentId" binding:"required"`
	MenuName  string `json:"menuName" binding:"required"`
	MenuCode  string `json:"menuCode" binding:"required"`
	Sort      int8   `json:"sort"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Query     string `json:"query"`
	Frame     string `json:"frame"`
	Cached    string `json:"cached"`
	Type      string `json:"type" binding:"required"`
	Hidden    string `json:"hidden"`
	Status    string `json:"status"`
	Perms     string `json:"perms"`
	Remark    string `json:"remark"`
}

func (m *SysMenuUpdateDTO) Convert(t *system.SysMenu) error {
	pId, err := strconv.ParseInt(m.ParentId, 10, 64)
	if err != nil {
		return err
	}

	mId, err := strconv.ParseInt(m.MenuId, 10, 64)
	if err != nil {
		return err
	}

	t.MenuId = mId
	t.ParentId = pId
	t.MenuCode = m.MenuCode
	t.MenuName = m.MenuName
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

type SysMenuQuery struct {
	MenuName string `form:"menuName" json:"menuName"`
	MenuCode string `form:"menuCode" json:"menuCode"`
	Status   string `form:"status" json:"status"`
}
