package system

import (
	"backend/model/common/request"
	"backend/model/system"
)

type SysDictTypeCreateDTO struct {
	DictName string `json:"dictName" binding:"required"`
	DictType string `json:"dictType" binding:"required"`
	Status   string `json:"status" binding:"required"`
	Remark   string `json:"remark" binding:"required"`
}

func (m *SysDictTypeCreateDTO) Convert(t *system.SysDictType) {
	t.DictName = m.DictName
	t.DictType = m.DictType
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysDictTypeUpdateDTO struct {
	DictTypeId int64  `json:"dictTypeId" binding:"required"`
	DictName   string `json:"dictName" binding:"required"`
	DictType   string `json:"dictType" binding:"required"`
	Status     string `json:"status" binding:"required"`
	Remark     string `json:"remark"`
}

func (m *SysDictTypeUpdateDTO) Convert(t *system.SysDictType) {
	t.DictTypeId = m.DictTypeId
	t.DictName = m.DictName
	t.DictType = m.DictType
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysDictTypeQueryDTO struct {
	PageInfo  request.PageInfo
	DictName  string `form:"dictName" json:"dictName"`
	DictType  string `form:"dictType" json:"dictType"`
	Status    string `form:"status" json:"status"`
	BeginTime string `form:"beginTime" json:"beginTime"`
	EndTime   string `form:"endTime" json:"endTime"`
}
