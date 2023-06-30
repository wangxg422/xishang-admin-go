package dto

import "backend/model/system"

type SysCreateDeptDTO struct {
	DeptId    int64  `json:"deptId,omitempty"`
	DeptName  string `json:"deptName,omitempty" binding:"required"`
	ParentId  int64  `json:"parentId,omitempty" binding:"required"`
	Ancestors string `json:"ancestors,omitempty"`
	OrderNum  int8   `json:"order_num,omitempty"`
	Leader    string `json:"leader,omitempty"`
	Phone     string `json:"phone,omitempty" binding:"required"`
	Email     string `json:"email,omitempty" binding:"required.email"`
	Status    int8   `json:"status,omitempty"`
	DelFlag   int8   `json:"delFlag,omitempty"`
	CreateBy  string `json:"createBy,omitempty"`
	UpdateBy  string `json:"updateBy,omitempty"`
}

func (m *SysCreateDeptDTO) Convert(t *system.SysDept) {
	t.DeptName = m.DeptName
	t.ParentId = m.ParentId
	t.Ancestors = m.Ancestors
	t.OrderNum = m.OrderNum
	t.Leader = m.Leader
	t.Phone = m.Phone
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}

type SysUpdateDeptDTO struct {
	DeptId    int64  `json:"deptId,omitempty" binding:"required"`
	DeptName  string `json:"deptName,omitempty"`
	ParentId  int64  `json:"parentId,omitempty"`
	Ancestors string `json:"ancestors,omitempty"`
	OrderNum  int8   `json:"order_num,omitempty"`
	Leader    string `json:"leader,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Status    int8   `json:"status,omitempty"`
	DelFlag   int8   `json:"delFlag,omitempty"`
	CreateBy  string `json:"createBy,omitempty"`
	UpdateBy  string `json:"updateBy,omitempty"`
}

func (m *SysUpdateDeptDTO) Convert(t *system.SysDept) {
	t.DeptName = m.DeptName
	t.ParentId = m.ParentId
	t.Ancestors = m.Ancestors
	t.OrderNum = m.OrderNum
	t.Leader = m.Leader
	t.Phone = m.Phone
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}
