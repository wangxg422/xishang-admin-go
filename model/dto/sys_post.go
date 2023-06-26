package dto

import (
	"backend/model/system"
	"time"
)

type SysCreatePostDTO struct {
	PostId     int64     `json:"postId,omitempty"`
	PostCode   string    `json:"postCode,omitempty" binding:"required"`
	PostName   string    `json:"postName,omitempty" binding:"required"`
	PostSort   int       ` json:"postSort,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysCreatePostDTO) Convert(t *system.SysPost) {
	t.PostId = m.PostId
	t.PostCode = m.PostCode
	t.PostName = m.PostName
	t.PostSort = m.PostSort
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}

type SysUpdatePostDTO struct {
	PostId     int64     `json:"postId,omitempty"`
	PostCode   string    `json:"postCode,omitempty" binding:"required"`
	PostName   string    `json:"postName,omitempty" binding:"required"`
	PostSort   int       ` json:"postSort,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	CreateBy   string    `json:"createBy,omitempty"`
	UpdateBy   string    `json:"updateBy,omitempty"`
	Remark     string    `json:"remark,omitempty"`
}

func (m *SysUpdatePostDTO) Convert(t *system.SysPost) {
	t.PostId = m.PostId
	t.PostCode = m.PostCode
	t.PostName = m.PostName
	t.PostSort = m.PostSort
	t.Status = m.Status
	t.CreateTime = m.CreateTime
	t.UpdateTime = m.UpdateTime
	t.CreateBy = m.CreateBy
	t.UpdateBy = m.UpdateBy
	t.Remark = m.Remark
}
