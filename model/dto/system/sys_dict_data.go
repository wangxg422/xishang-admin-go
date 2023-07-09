package system

import (
	"backend/model/common/request"
	"backend/model/system"
)

type SysDictDataCreateDTO struct {
	DictType  string `json:"dictType" binding:"required"`
	DictLabel string `json:"dictLabel" binding:"required"`
	DictValue string `json:"dictValue" binding:"required"`
	DictSort  int8   `json:"dictSort" binding:"required"`
	CssClass  string `json:"cssClass"`
	ListClass string `json:"listClass"`
	IsDefault int8   `json:"isDefault"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}

func (m *SysDictDataCreateDTO) Convert(t *system.SysDictData) {
	t.DictType = m.DictType
	t.DictLabel = m.DictLabel
	t.DictValue = m.DictValue
	t.DictSort = m.DictSort
	t.CssClass = m.CssClass
	t.ListClass = m.ListClass
	t.IsDefault = m.IsDefault
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysDictDataUpdateDTO struct {
	DictDataId int64  `json:"dictDataId" binding:"required"`
	DictType   string `json:"dictType" binding:"required"`
	DictLabel  string `json:"dictLabel"`
	DictValue  string `json:"dictValue"`
	DictSort   int8   `json:"dictSort"`
	CssClass   string `json:"cssClass"`
	ListClass  string `json:"listClass"`
	IsDefault  int8   `json:"isDefault"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

func (m *SysDictDataUpdateDTO) Convert(t *system.SysDictData) {
	t.DictDataId = m.DictDataId
	t.DictType = m.DictType
	t.DictLabel = m.DictLabel
	t.DictValue = m.DictValue
	t.DictSort = m.DictSort
	t.CssClass = m.CssClass
	t.ListClass = m.ListClass
	t.IsDefault = m.IsDefault
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysDictDataQueryDTO struct {
	PageInfo  request.PageInfo
	DictType  string `form:"dictType" json:"dictType" binding:"required"`
	DictLabel string `form:"dictLabel" json:"dictLabel"`
	Status    string `form:"status" json:"status"`
}
