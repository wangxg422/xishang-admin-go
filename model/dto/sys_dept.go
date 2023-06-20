package dto

import "backend/model/system"

type SysCreateDeptDTO struct {
	DeptId    int64  `json:"deptId,omitempty"`
	DeptName  string `json:"deptName,omitempty" binding:"required"`
	ParentId  int64  `json:"parentId,omitempty" binding:"required"`
	Ancestors int64  `json:"ancestors,omitempty"`
	OrderNum  int8   `json:"order_num,omitempty"`
	Leader    string `json:"leader,omitempty"`
	Phone     string `json:"phone,omitempty" binding:"required"`
	Email     string `json:"email,omitempty" binding:"required.email"`
	Status    string `json:"status,omitempty"`
	DelFlag   string `json:"delFlag,omitempty"`
	CreateBy  string `json:"createBy,omitempty"`
	UpdateBy  string `json:"updateBy,omitempty"`
}

func (d *SysCreateDeptDTO) Convert(user *system.SysDept) {
	user.DeptName = d.DeptName
	user.ParentId = d.ParentId
	user.Ancestors = d.Ancestors
	user.OrderNum = d.OrderNum
	user.Leader = d.Leader
	user.Phone = d.Phone
	user.Email = d.Email
	user.Status = d.Status
	user.DelFlag = d.DelFlag
	user.CreateBy = d.CreateBy
	user.UpdateBy = d.UpdateBy
}
