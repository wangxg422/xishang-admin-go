package system

import (
	"backend/model/system"
	"strconv"
)

type SysDeptCreateDTO struct {
	DeptName    string `json:"deptName" binding:"required"`
	DeptCode    string `json:"deptCode" binding:"required"`
	ParentId    string `json:"parentId" binding:"required"`
	Ancestors   string `json:"ancestors"`
	Sort        int8   `json:"sort"`
	Leader      string `json:"leader"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DelFlag     int8   `json:"delFlag"`
	CreateBy    string `json:"createBy"`
	UpdateBy    string `json:"updateBy"`
}

func (m *SysDeptCreateDTO) Convert() (system.SysDept, error) {
	var t system.SysDept

	pId, err := strconv.ParseInt(m.ParentId, 10, 64)
	if err != nil {
		return t, err
	}

	t.DeptName = m.DeptName
	t.DeptCode = m.DeptCode
	t.ParentId = pId
	t.Ancestors = m.Ancestors
	t.Sort = m.Sort
	t.Leader = m.Leader
	t.PhoneNumber = m.PhoneNumber
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy

	return t, nil
}

type SysDeptUpdateDTO struct {
	DeptId      string `json:"deptId" binding:"required"`
	DeptName    string `json:"deptName"`
	DeptCode    string `json:"deptCode"`
	ParentId    string `json:"parentId"`
	Ancestors   string `json:"ancestors"`
	Sort        int8   `json:"Sort"`
	Leader      string `json:"leader"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DelFlag     int8   `json:"delFlag"`
	CreateBy    string `json:"createBy"`
	UpdateBy    string `json:"updateBy"`
}

func (m *SysDeptUpdateDTO) Convert() (system.SysDept, error) {
	var t system.SysDept

	id, err := strconv.ParseInt(m.DeptId, 10, 64)
	if err != nil {
		return t, err
	}

	pId, err := strconv.ParseInt(m.ParentId, 10, 64)
	if err != nil {
		return t, err
	}

	t.DeptId = id
	t.DeptName = m.DeptName
	t.DeptCode = m.DeptCode
	t.ParentId = pId
	t.Ancestors = m.Ancestors
	t.Sort = m.Sort
	t.Leader = m.Leader
	t.PhoneNumber = m.PhoneNumber
	t.Email = m.Email
	t.Status = m.Status
	t.DelFlag = m.DelFlag
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy

	return t, nil
}

type SysDeptQueryDTO struct {
	DeptName string `form:"deptName" json:"deptName"`
	Status   string `form:"status" json:"status"`
}
