package system

import (
	sysModel "backend/model/system"
	"strconv"
)

type SysModuleCreateDTO struct {
	ModuleName string `json:"moduleName"`
	ModuleCode string `json:"moduleCode" binding:"required"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

func (m SysModuleCreateDTO) Convert(t *sysModel.SysModule) {
	t.ModuleName = m.ModuleName
	t.ModuleCode = m.ModuleCode
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysModuleUpdateDTO struct {
	ModuleId   string `json:"moduleId" binding:"required"`
	ModuleName string `json:"moduleName"`
	ModuleCode string `json:"moduleCode" binding:"required"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

func (m SysModuleUpdateDTO) Convert(t *sysModel.SysModule) error {
	id, err := strconv.ParseInt(m.ModuleId, 10, 64)
	if err != nil {
		return err
	}

	t.ModuleId = id
	t.ModuleName = m.ModuleName
	t.ModuleCode = m.ModuleCode
	t.Status = m.Status
	t.Remark = m.Remark

	return nil
}
