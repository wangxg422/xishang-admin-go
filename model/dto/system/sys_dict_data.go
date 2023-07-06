package system

import (
	"backend/model/system"
	"time"
)

type SysDictDataCreateDTO struct {
	DictDataId int64     `json:"dictDataId,omitempty"`
	DictSort   int8      `json:"dictSort,omitempty"`
	DictLabel  string    `json:"dictLabel,omitempty"`
	DictValue  string    `json:"dictValue,omitempty"`
	DictType   string    `json:"dictType,omitempty"`
	CssClass   string    `json:"cssClass,omitempty"`
	ListClass  string    `json:"listClass,omitempty"`
	IsDefault  int8      `json:"isDefault,omitempty"`
	Status     int8      `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysDictDataCreateDTO) SysDictTypeUpdateDTO(t *system.SysDictData) {
	t.DictDataId = m.DictDataId
	t.DictSort = m.DictSort
	t.DictLabel = m.DictLabel
	t.DictValue = m.DictValue
	t.DictType = m.DictType
	t.CssClass = m.CssClass
	t.ListClass = m.ListClass
	t.IsDefault = m.IsDefault
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysDictDataUpdateDTO struct {
	DictDataId int64     `json:"dictDataId,omitempty"`
	DictSort   int8      `json:"dictSort,omitempty"`
	DictLabel  string    `json:"dictLabel,omitempty"`
	DictValue  string    `json:"dictValue,omitempty"`
	DictType   string    `json:"dictType,omitempty"`
	CssClass   string    `json:"cssClass,omitempty"`
	ListClass  string    `json:"listClass,omitempty"`
	IsDefault  int8      `json:"isDefault,omitempty"`
	Status     int8      `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysDictDataUpdateDTO) SysDictDataUpdateDTO(t *system.SysDictData) {
	t.DictDataId = m.DictDataId
	t.DictSort = m.DictSort
	t.DictLabel = m.DictLabel
	t.DictValue = m.DictValue
	t.DictType = m.DictType
	t.CssClass = m.CssClass
	t.ListClass = m.ListClass
	t.IsDefault = m.IsDefault
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}
