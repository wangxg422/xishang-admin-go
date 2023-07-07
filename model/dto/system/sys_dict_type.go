package system

import (
	"backend/model/common/request"
	"backend/model/system"
	"time"
)

type SysDictTypeCreateDTO struct {
	DictTypeId int64     `json:"dictTypeId,omitempty"`
	DictName   string    `json:"dictName,omitempty"`
	DictType   string    `json:"dictType,omitempty"`
	Status     int8      `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysDictTypeCreateDTO) Convert(t *system.SysDictType) {
	t.DictTypeId = m.DictTypeId
	t.DictName = m.DictName
	t.DictType = m.DictType
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysDictTypeUpdateDTO struct {
	DictTypeId int64     `json:"dictTypeId,omitempty"`
	DictName   string    `json:"dictName,omitempty"`
	DictType   string    `json:"dictType,omitempty"`
	Status     int8      `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysDictTypeUpdateDTO) Convert(t *system.SysDictType) {
	t.DictTypeId = m.DictTypeId
	t.DictName = m.DictName
	t.DictType = m.DictType
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysDictTypeQueryDTO struct {
	PageInfo  request.PageInfo
	DictName  string `form:"dictName" json:"dictName"`
	DictType  string `form:"dictType" json:"dictType"`
	Status    int8   `form:"status" json:"status"`
	BeginTime string `form:"beginTime" json:"beginTime"`
	EndTime   string `form:"endTime" json:"endTime"`
}
