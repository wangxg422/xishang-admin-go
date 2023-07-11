package system

import (
	"backend/model/common/request"
	"backend/model/system"
	"time"
)

type SysCreateUserDTO struct {
	DeptId      int64   `json:"deptId,omitempty"`
	UserName    string  `json:"userName,omitempty" binding:"required"`
	NickName    string  `json:"nickName,omitempty" binding:"required"`
	Email       string  `json:"email,omitempty"`
	PhoneNumber string  `json:"phoneNumber,omitempty"`
	Sex         string  `json:"sex"`
	Password    string  `json:"password,omitempty" binding:"required"`
	Status      string  `json:"status,omitempty"`
	Remark      string  `json:"remark,omitempty"`
	RoleIds     []int64 `json:"roleIds,omitempty"`
	PostIds     []int64 `json:"postIds,omitempty"`
}

func (m *SysCreateUserDTO) Convert(user *system.SysUser) {
	user.DeptId = m.DeptId
	user.UserName = m.UserName
	user.NickName = m.NickName
	user.Email = m.Email
	user.PhoneNumber = m.PhoneNumber
	user.Sex = m.Sex
	user.Password = m.Password
	user.Status = m.Status
	user.Remark = m.Remark
}

type SysUpdateUserDTO struct {
	UserId      int64     `json:"userId,omitempty" binding:"required"`
	DeptId      int64     `json:"deptId,omitempty"`
	UserName    string    `json:"userName,omitempty"`
	NickName    string    `json:"nickName,omitempty"`
	UserType    string    `json:"userType,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Sex         string    `json:"sex"`
	Avatar      string    `json:"avatar,omitempty"`
	Password    string    `json:"password,omitempty"`
	Status      string    `json:"status"`
	DelFlag     int8      `json:"delFlag"`
	CreateTime  time.Time `json:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty"`
	CreateBy    string    `json:"createBy,omitempty"`
	UpdateBy    string    `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
}

func (m *SysUpdateUserDTO) Convert(user *system.SysUser) {
	user.UserId = m.UserId
	user.DeptId = m.DeptId
	user.UserName = m.UserName
	user.NickName = m.NickName
	user.UserType = m.UserType
	user.Email = m.Email
	user.PhoneNumber = m.PhoneNumber
	user.Sex = m.Sex
	user.Avatar = m.Avatar
	user.Password = m.Password
	user.Status = m.Status
	user.DelFlag = m.DelFlag
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}

type SysUserQueryDTO struct {
	request.PageInfo
	UserName    string `form:"userName" json:"userName"`
	DeptId      int64  `form:"deptId" json:"deptId"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber"`
	Status      string `form:"status" json:"status"`
	BeginTime   string `form:"beginTime" json:"beginTime"`
	EndTime     string `form:"endTime" json:"endTime"`
}
