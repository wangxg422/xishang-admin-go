package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	sysVo "backend/model/vo/system"
	"backend/utils"
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
	utils.ConcatOneEqualsInt8WhereCondition(db, "status", params.Status)

	var types []sysModel.SysDictType
	res := db.Find(&types)

	pageResult.Total = total
	pageResult.Rows = types

	return pageResult, res.Error
}
