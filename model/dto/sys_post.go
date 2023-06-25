package dto

import (
	"backend/model/system"
	"time"
)

type SysCreatePostDTO struct {
	PostId     int64     `json:"postid,omitempty"`
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

func (m *SysCreatePostDTO) Convert(user *system.SysPost) {
	user.PostId = m.PostId
	user.PostCode = m.PostCode
	user.PostName = m.PostName
	user.PostSort = m.PostSort
	user.Status = m.Status
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}

type SysUpdatePostDTO struct {
	PostId     int64     `json:"postid,omitempty"`
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

func (m *SysUpdatePostDTO) Convert(user *system.SysPost) {
	user.PostId = m.PostId
	user.PostCode = m.PostCode
	user.PostName = m.PostName
	user.PostSort = m.PostSort
	user.Status = m.Status
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}
