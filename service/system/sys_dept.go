package system

import (
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
)

type SysDeptService struct {
}

func (m *SysDeptService) CreateDept(d *sysModel.SysDept) error {
	res := global.DB.Create(&d)

	return res.Error
}

func (m *SysDeptService) UpdateDept(d *sysModel.SysDept) error {
	res := global.DB.Model(&sysModel.SysDept{DeptId: d.DeptId}).Updates(&d)
	return res.Error
}

func (m *SysDeptService) DeleteDept(id int64) error {
	res := global.DB.Model(&sysModel.SysDept{DeptId: id}).Update("del_flag", enmu.DelFlagDeleted.Value())
	return res.Error
}

func (m *SysDeptService) GetDeptById(id int64) (sysModel.SysDept, error) {
	user := sysModel.SysDept{
		DeptId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.StatusNormal.Value())
	return user, res.Error
}

func (m *SysDeptService) GetDeptByIds(ids []string) ([]sysModel.SysDept, error) {
	var depts []sysModel.SysDept
	res := global.DB.
		Where("dept_id IN ? and status = ?", ids, enmu.StatusNormal.Value()).
		Order("dept_id").
		Find(&depts)
	return depts, res.Error
}

func (m *SysDeptService) GetAllDept() ([]sysModel.SysDept, error) {
	var depts []sysModel.SysDept
	res := global.DB.
		Where("status = ?", enmu.StatusNormal.Value()).
		Order("dept_id").
		Find(&depts)
	return depts, res.Error
}

func (m *SysDeptService) GetDept(params *sysDto.SysDeptQueryDTO) ([]sysModel.SysDept, error) {
	db := global.DB.Model(&sysModel.SysDept{})

	utils.ConcatOneLikeWhereCondition(db, "dept_name", params.DeptName)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)

	var depts []sysModel.SysDept
	err := db.Order("parent_id, sort").Find(&depts).Error

	return depts, err
}
