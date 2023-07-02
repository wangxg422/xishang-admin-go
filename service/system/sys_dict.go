package system

import (
	"backend/common/enmu"
	"backend/global"
	"backend/model/system"
)

type SysDictService struct {
}

func (m *SysDictService) GetDictDataByType(dictType string) ([]system.SysDictData, error) {
	var data []system.SysDictData
	res := global.DB.
		Where("dict_type = ? and status = ?", dictType, enmu.StatusNormal.Value()).
		Order("dict_sort").
		Find(&data)
	return data, res.Error
}
