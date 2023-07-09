package system

import (
	"backend/model/common/request"
	"backend/model/system"
)

type SysCreatePostDTO struct {
	PostCode string `json:"postCode" binding:"required"`
	PostName string `json:"postName" binding:"required"`
	PostSort int    ` json:"postSort" binding:"required"`
	Status   string `json:"status" binding:"required"`
	CreateBy string `json:"createBy"`
	Remark   string `json:"remark"`
}

func (m *SysCreatePostDTO) Convert(t *system.SysPost) {
	t.PostCode = m.PostCode
	t.PostName = m.PostName
	t.PostSort = m.PostSort
	t.Status = m.Status
	t.CreateBy = m.CreateBy
	t.Remark = m.Remark
}

type SysUpdatePostDTO struct {
	PostId   int64  `json:"postId,omitempty"`
	PostCode string `json:"postCode,omitempty" binding:"required"`
	PostName string `json:"postName,omitempty" binding:"required"`
	PostSort int    ` json:"postSort,omitempty" binding:"required"`
	Status   string `json:"status,omitempty" binding:"required"`
	Remark   string `json:"remark,omitempty"`
}

func (m *SysUpdatePostDTO) Convert(t *system.SysPost) {
	t.PostId = m.PostId
	t.PostCode = m.PostCode
	t.PostName = m.PostName
	t.PostSort = m.PostSort
	t.Status = m.Status
	t.Remark = m.Remark
}

type SysPostQueryDTO struct {
	PageInfo request.PageInfo
	PostCode string `form:"postCode" json:"postCode"`
	PostName string `form:"postName" json:"postName"`
	Status   string `form:"status" json:"status"`
}
