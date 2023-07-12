package system

import (
	"backend/model/common/request"
	"backend/model/system"
	"time"
)

type SysUserCreateDTO struct {
	DeptId      int64  `json:"deptId"`
	UserName    string `json:"userName" binding:"required"`
	UserNumber  string `json:"userNumber" binding:"required"`
	RealName    string `json:"realName" binding:"required"`
	NickName    string `json:"nickName" binding:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	UserType    string `json:"userType"`
	Sex         string `json:"sex"`
	Password    string `json:"password" binding:"required"`
	Status      string `json:"status"`
	Remark      string `json:"remark"`

	RoleIds []int64 `json:"roleIds,omitempty"`
	PostIds []int64 `json:"postIds,omitempty"`
}

func (m *SysUserCreateDTO) Convert(user *system.SysUser) {
	user.DeptId = m.DeptId
	user.UserName = m.UserName
	user.UserNumber = m.UserNumber
	user.RealName = m.RealName
	user.NickName = m.NickName
	user.Email = m.Email
	user.PhoneNumber = m.PhoneNumber
	user.UserType = m.UserType
	user.Sex = m.Sex
	user.Password = m.Password
	user.Status = m.Status
	user.Remark = m.Remark
}

type SysUserUpdateDTO struct {
	UserId      int64     `json:"userId" binding:"required"`
	DeptId      int64     `json:"deptId"`
	UserName    string    `json:"userName" binding:"required"`
	UserNumber  string    `json:"userNumber" binding:"required"`
	RealName    string    `json:"realName" binding:"required"`
	NickName    string    `json:"nickName"`
	UserType    string    `json:"userType"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Sex         string    `json:"sex"`
	Avatar      string    `json:"avatar"`
	Password    string    `json:"password" binding:"required"`
	Status      string    `json:"status"`
	UpdateTime  time.Time `json:"updateTime"`
	UpdateBy    string    `json:"updateBy"`
	Remark      string    `json:"remark"`

	RoleIds []int64 `json:"roleIds"`
	PostIds []int64 `json:"postIds"`
}

func (m *SysUserUpdateDTO) Convert(user *system.SysUser) {
	user.UserId = m.UserId
	user.DeptId = m.DeptId
	user.UserName = m.UserName
	user.UserNumber = m.UserNumber
	user.RealName = m.RealName
	user.NickName = m.NickName
	user.UserType = m.UserType
	user.Email = m.Email
	user.PhoneNumber = m.PhoneNumber
	user.Sex = m.Sex
	user.Avatar = m.Avatar
	user.Password = m.Password
	user.Status = m.Status
	user.UpdateTime = m.UpdateTime
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}

type SysUserQueryDTO struct {
	request.PageInfo
	UserName    string `form:"userName" json:"userName"`
	UserUmber   string `form:"userNumber" json:"userNumber"`
	RealName    string `form:"realName" json:"realName"`
	DeptId      int64  `form:"deptId" json:"deptId"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber"`
	Status      string `form:"status" json:"status"`
	BeginTime   string `form:"beginTime" json:"beginTime"`
	EndTime     string `form:"endTime" json:"endTime"`
}
