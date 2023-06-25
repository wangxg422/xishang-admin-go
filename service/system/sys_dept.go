package system

import (
	"backend/common/enmu"
	"backend/global"
	sysModel "backend/model/system"
)

type SysDeptService struct {
}

func (u *SysDeptService) CreateDept(d *sysModel.SysDept) error {
	res := global.DB.Create(&d)

	return res.Error
}

func (u *SysDeptService) UpdateUser(d *sysModel.SysDept) error {
	res := global.DB.Model(&sysModel.SysDept{DeptId: d.DeptId}).Updates(&d)
	return res.Error
}

func (u *SysDeptService) DeleteDept(id int64) error {
	res := global.DB.Model(&sysModel.SysDept{DeptId: id}).Update("del_flag", enmu.EnmuGroupApp.DelFlagDelete.GetCode())
	return res.Error
}

func (u *SysDeptService) ListUser() ([]sysModel.SysDept, error) {
	var list []sysModel.SysDept

	res := global.DB.Where("del_flag = ?", enmu.EnmuGroupApp.DelFlagNormal.GetCode()).Find(&list)

	return list, res.Error
}

func (u *SysDeptService) GetDeptById(id int64) (sysModel.SysDept, error) {
	user := sysModel.SysDept{
		DeptId: id,
	}

	res := global.DB.Take(&user, id).Where("del_flag = ?", enmu.EnmuGroupApp.StatusNormal.GetCode())
	return user, res.Error
}
