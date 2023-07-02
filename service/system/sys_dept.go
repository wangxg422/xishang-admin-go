package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
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

func (m *SysDeptService) ListDept() ([]sysModel.SysDept, error) {
	var list []sysModel.SysDept

	res := global.DB.Where("del_flag = ?", enmu.DelFlagNormal.Value()).Find(&list)

	return list, res.Error
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
