package dto

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
	Status            string    `json:"status,omitempty"`
	DelFlag           string    `json:"delFlag,omitempty"`
	CreateTime        time.Time `json:"createTime,omitempty"`
	UpdateTime        time.Time `json:"updateTime,omitempty"`
	CreateBy          string    `json:"createBy,omitempty"`
	UpdateBy          string    `json:"updateBy,omitempty"`
	Remark            string    `json:"remark,omitempty"`
}

func (m *SysCreateRoleDTO) Convert(user *system.SysRole) {
	user.RoleId = m.RoleId
	user.RoleName = m.RoleName
	user.RoleKey = m.RoleKey
	user.RoleSort = m.RoleSort
	user.DataScope = m.DataScope
	user.MenuCheckStrictly = m.MenuCheckStrictly
	user.DeptCheckStrictly = m.DeptCheckStrictly
	user.Status = m.Status
	user.DelFlag = m.DelFlag
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
}

type SysUpdateRoleDTO struct {
	RoleId            int64     `json:"roleId,omitempty"`
	RoleName          string    `json:"roleName,omitempty" binding:"required"`
	RoleKey           string    `json:"roleKey,omitempty"`
	RoleSort          int64     `json:"roleSort,omitempty"`
	DataScope         int8      `json:"dataScope,omitempty"`
	MenuCheckStrictly string    `json:"menuCheckStrictly,omitempty"`
	DeptCheckStrictly string    `json:"deptCheckStrictly,omitempty"`
	Status            string    `json:"status,omitempty"`
	DelFlag           string    `json:"delFlag,omitempty"`
	CreateTime        time.Time `json:"createTime,omitempty"`
	UpdateTime        time.Time `json:"updateTime,omitempty"`
	CreateBy          string    `json:"createBy,omitempty"`
	UpdateBy          string    `json:"updateBy,omitempty"`
	Remark            string    `json:"remark,omitempty"`
}

func (m *SysUpdateRoleDTO) Convert(user *system.SysRole) {
	user.RoleId = m.RoleId
	user.RoleName = m.RoleName
	user.RoleKey = m.RoleKey
	user.RoleSort = m.RoleSort
	user.DataScope = m.DataScope
	user.MenuCheckStrictly = m.MenuCheckStrictly
	user.DeptCheckStrictly = m.DeptCheckStrictly
	user.Status = m.Status
	user.DelFlag = m.DelFlag
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
}
