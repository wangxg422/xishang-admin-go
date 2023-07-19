package system

import (
	"backend/model/common/request"
	"backend/model/system"
	"strconv"
)

type SysRoleCreateDTO struct {
	RoleName          string `json:"roleName" binding:"required"`
	RoleCode          string `json:"roleCode" binding:"required"`
	RolePerms         string `json:"rolePerms"`
	Sort              int64  `json:"roleSort"`
	DataScope         int8   `json:"dataScope"`
	MenuCheckStrictly int8   `json:"menuCheckStrictly"`
	DeptCheckStrictly int8   `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
}

func (m *SysRoleCreateDTO) Convert() system.SysRole {
	var t system.SysRole

	t.RoleName = m.RoleName
	t.RoleCode = m.RoleCode
	t.RolePerms = m.RolePerms
	t.Sort = m.Sort
	t.DataScope = m.DataScope
	t.MenuCheckStrictly = m.MenuCheckStrictly
	t.DeptCheckStrictly = m.DeptCheckStrictly
	t.Status = m.Status
	t.Remark = m.Remark

	return t
}

type SysRoleUpdateDTO struct {
	RoleId            string `json:"roleId"`
	RoleName          string `json:"roleName" binding:"required"`
	RoleCode          string `json:"roleCode" binding:"required"`
	RolePerms         string `json:"rolePerms"`
	Sort              int64  `json:"sort"`
	DataScope         int8   `json:"dataScope"`
	MenuCheckStrictly int8   `json:"menuCheckStrictly"`
	DeptCheckStrictly int8   `json:"deptCheckStrictly"`
	Status            string `json:"status"`
	Remark            string `json:"remark"`
}

func (m *SysRoleUpdateDTO) Convert() (system.SysRole, error) {
	var t system.SysRole
	id, err := strconv.ParseInt(m.RoleId, 10, 64)
	if err != nil {
		return t, err
	}

	t.RoleId = id
	t.RoleName = m.RoleName
	t.RoleCode = m.RoleCode
	t.RolePerms = m.RolePerms
	t.Sort = m.Sort
	t.DataScope = m.DataScope
	t.MenuCheckStrictly = m.MenuCheckStrictly
	t.DeptCheckStrictly = m.DeptCheckStrictly
	t.Status = m.Status
	t.Remark = m.Remark

	return t, nil
}

type SysRoleQueryDTO struct {
	request.PageInfo
	RoleName  string `form:"roleName" json:"roleName"`
	RoleCode  string `form:"roleCode" json:"roleCode"`
	Status    string `form:"status" json:"status"`
	RolePerms string `form:"rolePerms" json:"rolePerms"`
	BeginTime string `form:"beginTime" json:"beginTime"`
	EndTime   string `form:"endTime" json:"endTime"`
}
