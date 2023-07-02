package system

import "time"

type SysDictData struct {
	DictCode   int64     `gorm:"primaryKey;column:dict_code" json:"dictCode,omitempty"`
	DictSort   int8      `gorm:"column:dict_sort" json:"dict_sort,omitempty"`
	DictLabel  string    `gorm:"primaryKey;column:dict_label" json:"dictLabel,omitempty"`
	DictValue  string    `gorm:"primaryKey;column:dict_value" json:"dictValue,omitempty"`
	DictType   string    `gorm:"primaryKey;column:dict_type" json:"dictType,omitempty"`
	CssClass   string    `gorm:"primaryKey;column:css_class" json:"cssClass,omitempty"`
	ListClass  string    `gorm:"primaryKey;column:list_class" json:"listClass,omitempty"`
	IsDefault  string    `gorm:"column:is_default;default:0" json:"isDefault,omitempty"`
	Status     int8      `gorm:"column:status;default:0" json:"status,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark" json:"remark,omitempty"`
}
