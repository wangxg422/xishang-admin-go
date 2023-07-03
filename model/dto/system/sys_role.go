package system

import (
	"backend/model/system"
	"time"
)

type SysCreateRoleDTO struct {
	RoleId            int64     `json:"roleId,omitempty"`
	RoleName          string    `json:"roleName,omitempty" binding:"required"`
	RoleKey           string    `json:"roleKey,omitempty"`
	RoleSort          int64     `json:"roleSort,omitempty"`
	DataScope         int8      `json:"dataScope,omitempty"`
	MenuCheckStrictly string    `json:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly string    `json:"deptCheckStrictly,omitempty"`
	Status            int8      `json:"status,omitempty"`
	DelFlag           int8      `json:"delFlag,omitempty"`
	CreateTime        time.Time `json:"createTime,omitempty"`
	UpdateTime        time.Time `json:"updateTime,omitempty"`
	CreateBy          string    `json:"createBy,omitempty"`
	UpdateBy          string    `json:"updateBy,omitempty"`
	Remark            string    `json:"remark,omitempty"`
}

func (m *SysCreateRoleDTO) Convert(t *system.SysRole) {
	t.RoleId = m.RoleId
	t.RoleName = m.RoleName
	t.RoleKey = m.RoleKey
	t.RoleSort = m.RoleSort
	t.DataScope = m.DataScope
	t.MenuCheckStrictly = m.MenuCheckStrictly
	t.DeptCheckStrictly = m.DeptCheckStrictly
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}

type SysUpdateRoleDTO struct {
	RoleId            int64     `json:"roleId,omitempty"`
	RoleName          string    `json:"roleName,omitempty" binding:"required"`
	RoleKey           string    `json:"roleKey,omitempty"`
	RoleSort          int64     `json:"roleSort,omitempty"`
	DataScope         int8      `json:"dataScope,omitempty"`
	MenuCheckStrictly string    `json:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly string    `json:"deptCheckStrictly,omitempty"`
	Status            int8      `json:"status,omitempty"`
	DelFlag           int8      `json:"delFlag,omitempty"`
	CreateTime        time.Time `json:"createTime,omitempty"`
	UpdateTime        time.Time `json:"updateTime,omitempty"`
	CreateBy          string    `json:"createBy,omitempty"`
	UpdateBy          string    `json:"updateBy,omitempty"`
	Remark            string    `json:"remark,omitempty"`
}

func (m *SysUpdateRoleDTO) Convert(t *system.SysRole) {
	t.RoleId = m.RoleId
	t.RoleName = m.RoleName
	t.RoleKey = m.RoleKey
	t.RoleSort = m.RoleSort
	t.DataScope = m.DataScope
	t.MenuCheckStrictly = m.MenuCheckStrictly
	t.DeptCheckStrictly = m.DeptCheckStrictly
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}
