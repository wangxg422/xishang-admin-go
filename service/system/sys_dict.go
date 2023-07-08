package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
	"errors"
)

type SysDictService struct {
}

func (m *SysDictService) GetDictDataByType(dictType string) ([]sysModel.SysDictData, error) {
	var data []sysModel.SysDictData
	res := global.DB.
		Where("dict_type = ? and status = ?", dictType, enmu.StatusNormal.Value()).
		Order("dict_sort").
		Find(&data)
	return data, res.Error
}

func (m *SysDictService) GetDictTypePage(params *sysDto.SysDictTypeQueryDTO) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysDictType{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	likeArr := []string{
		"dict_name",
		"dict_type",
	}
	utils.ConcatLikeWhereCondition(db, likeArr, params.DictName, params.DictType)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)

	var types []sysModel.SysDictType
	res := db.Find(&types)

	pageResult.Total = total
	pageResult.Rows = types

	return pageResult, res.Error
}

func (m *SysDictService) CreateDictType(dictType *sysModel.SysDictType) error {
	var existType sysModel.SysDictType
	err := global.DB.
		Where("dict_type = ?", dictType.DictType).
		Find(&existType).Limit(1).Error
	if err != nil {
		return err
	}

	if existType.DictTypeId != 0 {
		return errors.New("字典类型 " + dictType.DictType + " 已经存在")
	}

	return global.DB.Create(dictType).Error
}

func (m *SysDictService) UpdateDictType(dictType *sysModel.SysDictType) error {
	// 这里保证零值也能更新
	return global.DB.Select("dict_name", "dict_type", "status",
		"update_time", "update_by", "remark").
		Updates(dictType).Error
}

func (m *SysDictService) GetDictTypeById(typeId int64) (sysModel.SysDictType, error) {
	var dictType sysModel.SysDictType
	res := global.DB.
		Where("dict_type_id = ?", typeId).
		Find(&dictType)

	return dictType, res.Error
}

func (m *SysDictService) DeleteDictType(ids []int64) error {
	return global.DB.Where("dict_type_id IN ?", ids).Delete(&sysModel.SysDictType{}).Error
}
