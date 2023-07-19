package system

import (
	"backend/model/system"
)

type SysDeptCreateDTO struct {
	DeptId    int64  `json:"deptId"`
	DeptName  string `json:"deptName" binding:"required"`
	ParentId  int64  `json:"parentId" binding:"required"`
	Ancestors string `json:"ancestors"`
	Sort      int8   `json:"sort"`
	Leader    string `json:"leader"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required.email"`
	Status    string `json:"status"`
	DelFlag   int8   `json:"delFlag"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
}

func (m *SysDeptCreateDTO) Convert(t *system.SysDept) {
	t.DeptName = m.DeptName
	t.ParentId = m.ParentId
	t.Ancestors = m.Ancestors
	t.Sort = m.Sort
	t.Leader = m.Leader
	t.Phone = m.Phone
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}

type SysDeptUpdateDTO struct {
	DeptId    int64  `json:"deptId" binding:"required"`
	DeptName  string `json:"deptName"`
	ParentId  int64  `json:"parentId"`
	Ancestors string `json:"ancestors"`
	Sort      int8   `json:"Sort"`
	Leader    string `json:"leader"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	DelFlag   int8   `json:"delFlag"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
}

func (m *SysDeptUpdateDTO) Convert(t *system.SysDept) {
	t.DeptName = m.DeptName
	t.ParentId = m.ParentId
	t.Ancestors = m.Ancestors
	t.Sort = m.Sort
	t.Leader = m.Leader
	t.Phone = m.Phone
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
}

type SysDeptQueryDTO struct {
	DeptName string `form:"deptName" json:"deptName"`
	Status   string `form:"status" json:"status"`
}
