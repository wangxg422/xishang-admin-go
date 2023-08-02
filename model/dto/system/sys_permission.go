package system

import (
	sysModel "backend/model/system"
	"strconv"
)

type SysPermissionCreateDTO struct {
	ParentId   int64  `json:"parentId,string" binding:"required"`
	ModuleCode string `json:"moduleCode" binding:"required"`
	PermName   string `json:"permName"`
	PermCode   string `json:"postCode" binding:"required"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

func (m SysPermissionCreateDTO) Convert(t *sysModel.SysPermission) {
	t.ParentId = m.ParentId
	t.ModuleCode = m.ModuleCode
	t.PermName = m.PermName
	t.PermCode = m.PermCode
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysPermissionUpdateDTO struct {
	PermId     string `json:"permId" binding:"required"`
	ParentId   int64  `json:"parentId,string" binding:"required"`
	ModuleCode string `json:"moduleCode" binding:"required"`
	PermName   string `json:"permName"`
	PermCode   string `json:"postCode" binding:"required"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
}

func (m SysPermissionUpdateDTO) Convert(t *sysModel.SysPermission) error {
	id, err := strconv.ParseInt(m.PermId, 10, 64)
	if err != nil {
		return err
	}
	t.PermId = id

	t.ParentId = m.ParentId
	t.ModuleCode = m.ModuleCode
	t.PermName = m.PermName
	t.PermCode = m.PermCode
	t.Status = m.Status
	t.Remark = m.Remark

	return nil
}
