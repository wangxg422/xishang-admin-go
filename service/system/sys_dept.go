package system

import (
	"backend/common/constant"
	"backend/common/enmu"
	"backend/global"
	sysDto "backend/model/dto/system"
	sysModel "backend/model/system"
	"backend/utils"
	"errors"
)

type SysDeptService struct {
}

func (m *SysDeptService) CreateDept(d *sysModel.SysDept) error {
	// 检查DeptCode是否已经存在
	existDept, err := m.GetDeptByCode(d.DeptCode)
	if err != nil {
		return err
	}
	if existDept.DeptId != 0 {
		return errors.New("部门编码 " + d.DeptCode + " 已存在")
	}

	return global.DB.Create(&d).Error
}

func (m *SysDeptService) GetDeptByCode(code string) (sysModel.SysDept, error) {
	var dept sysModel.SysDept
	err := global.DB.Where("dept_code = ? AND del_flag = ?", code, enmu.DelFlagNormal.Value()).Find(&dept).Error
	return dept, err
}

func (m *SysDeptService) UpdateDept(d *sysModel.SysDept) error {
	if d.DeptId == 0 {
		return errors.New("deptId is null")
	}

	// 检查更新的dept_code是否已经存在
	existMenu, err := m.GetDeptByCode(d.DeptCode)
	if err != nil {
		return err
	}
	if existMenu.DeptId != 0 && existMenu.DeptId != d.DeptId {
		return errors.New("部门编码 " + d.DeptCode + " 已存在")
	}

	return global.DB.Omit(constant.UpdateOmit...).Save(d).Error
}

func (m *SysDeptService) DeleteDept(id int64) error {
	// 有子部门的禁止删除
	count, err := m.GetChildDeptCountById(id)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("含有下级部门，禁止删除")
	}

	// TODO 部门中存在用户的禁止删除

	return global.DB.Model(&sysModel.SysDept{DeptId: id}).Update("del_flag", enmu.DelFlagDeleted.Value()).Error
}

func (m *SysDeptService) GetChildDeptCountById(id int64) (int64, error) {
	var count int64 = 0
	err := global.DB.Model(&sysModel.SysDept{}).Where("parent_id = ? AND del_flag = ?", id, enmu.DelFlagNormal.Value()).Count(&count).Error
	return count, err
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
		Where("status = ? AND del_flag = ?", enmu.StatusNormal.Value(), enmu.DelFlagNormal.Value()).
		Order("dept_id").
		Find(&depts)
	return depts, res.Error
}

func (m *SysDeptService) GetDept(params *sysDto.SysDeptQueryDTO) ([]sysModel.SysDept, error) {
	db := global.DB.Model(&sysModel.SysDept{})

	utils.ConcatOneLikeWhereCondition(db, "dept_name", params.DeptName)
	utils.ConcatOneEqualsStrWhereCondition(db, "status", params.Status)
	utils.ConcatOneEqualsInt8WhereCondition(db, "del_flag", enmu.DelFlagNormal.Value())

	var depts []sysModel.SysDept
	err := db.Order("parent_id, sort").Find(&depts).Error

	return depts, err
}
