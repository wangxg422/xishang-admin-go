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

	likeArr := []string{
		"dict_name",
		"dict_type",
	}
	utils.ConcatLikeWhereCondition(db, likeArr, params.DictName, params.DictType)
	utils.ConcatTimeRangeWhereCondition(db, params.BeginTime, params.EndTime)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

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

func (m *SysDictService) UpdateDictType(data *sysModel.SysDictType) error {
	// 查询是否存在postCode的岗位
	existType, err := m.GetDictTypeByType(data.DictType)
	if err != nil {
		return err
	}

	if existType.DictTypeId != 0 && existType.DictTypeId != data.DictTypeId {
		return errors.New("职位编码 " + data.DictType + " 已经存在")
	}

	vMap, err := utils.StructToMap(data)
	if err != nil {
		return err
	}

	utils.DeleteKvWhenUpdate(vMap)

	return global.DB.Model(&sysModel.SysDictType{DictTypeId: data.DictTypeId}).Updates(vMap).Error
}

func (m *SysDictService) GetDictTypeById(typeId int64) (sysModel.SysDictType, error) {
	var dictType sysModel.SysDictType
	res := global.DB.
		Where("dict_type_id = ?", typeId).
		Find(&dictType)

	return dictType, res.Error
}

func (m *SysDictService) GetDictTypeByType(data string) (sysModel.SysDictType, error) {
	var dictType sysModel.SysDictType

	return dictType, global.DB.Where("dict_type = ?", data).Find(&dictType).Error
}

func (m *SysDictService) DeleteDictType(ids []int64) error {
	return global.DB.Where("dict_type_id IN ?", ids).Delete(&sysModel.SysDictType{}).Error
}

func (m *SysDictService) GetDictTypeAll() ([]sysModel.SysDictType, error) {
	var dictTypes []sysModel.SysDictType
	err := global.DB.Find(&dictTypes).Error

	return dictTypes, err
}

func (m *SysDictService) GetDictDataPage(params *sysDto.SysDictDataQueryDTO) (sysVo.PageResult, error) {
	pageResult := sysVo.PageResult{}

	db := global.DB.Model(&sysModel.SysDictData{})

	likeArr := []string{
		"dict_label",
	}
	equalsArr := []string{
		"dict_type",
		"status",
	}
	utils.ConcatLikeWhereCondition(db, likeArr, params.DictLabel)
	utils.ConcatEqualsStrWhereCondition(db, equalsArr, params.DictType, params.Status)
	db.Order("dict_sort")

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return pageResult, err
	}

	limit, offset := params.PageInfo.Paging()
	db = db.Limit(limit).Offset(offset)

	var types []sysModel.SysDictData
	res := db.Find(&types)

	pageResult.Total = total
	pageResult.Rows = types

	return pageResult, res.Error
}

func (m *SysDictService) CreateDictData(data *sysModel.SysDictData) error {
	// 根据dictType和dictValue检查数据是否已经存在
	existData, err := m.GetDictDataByTypeAndValue(data.DictType, data.DictValue)
	if err != nil {
		return err
	}
	if existData.DictDataId != 0 {
		return errors.New("字典数据已存在")
	}

	return global.DB.Create(data).Error
}

func (m *SysDictService) GetDictDataByTypeAndValue(dictType string, dictValue string) (sysModel.SysDictData, error) {
	var data sysModel.SysDictData
	res := global.DB.Where("dict_type = ? and dict_value = ?", dictType, dictValue).Limit(1).Find(&data)
	return data, res.Error
}

func (m *SysDictService) UpdateDictData(data *sysModel.SysDictData) error {
	// 先检查要更新的目标值是否存在
	existData, err := m.GetDictDataByTypeAndValue(data.DictType, data.DictValue)
	if err != nil {
		return err
	}

	if existData.DictDataId != 0 && existData.DictDataId != data.DictDataId {
		return errors.New("字典数据 " + data.DictType + ":" + data.DictValue + " 已存在")
	}

	vMap, err := utils.StructToMap(data)
	if err != nil {
		return err
	}

	utils.DeleteKvWhenUpdate(vMap)

	return global.DB.Model(&sysModel.SysDictData{DictDataId: data.DictDataId}).Updates(vMap).Error
}

func (m *SysDictService) GetDictDataById(id int64) (sysModel.SysDictData, error) {
	var dictData sysModel.SysDictData
	err := global.DB.Where("dict_data_id = ?", id).Find(&dictData).Error

	return dictData, err
}

func (m *SysDictService) DeleteDictData(ids []int64) error {
	return global.DB.Where("dict_data_id IN ?", ids).Delete(&sysModel.SysDictData{}).Error
}
