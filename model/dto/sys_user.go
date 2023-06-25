package dto

import (
	"backend/model/system"
	"time"
)

type SysCreateUserDTO struct {
	UserId      int64     `json:"userid,omitempty"`
	DeptId      int64     `json:"deptId,omitempty"`
	UserName    string    `json:"username,omitempty" binding:"required"`
	NickName    string    `json:"nickname,omitempty" binding:"required"`
	UserType    string    `json:"userType,omitempty"`
	Email       string    `json:"email,omitempty" binding:"required,email"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Sex         string    `json:"sex,omitempty" binding:"required"`
	Avatar      string    `json:"avatar,omitempty"`
	Password    string    `json:"password,omitempty" binding:"required"`
	Status      string    `json:"status,omitempty"`
	DelFlag     string    `json:"delFlag,omitempty"`
	LoginIp     string    `json:"loginIp,omitempty"`
	LoginDate   time.Time `json:"loginDate,omitempty"`
	CreateTime  time.Time `json:"createTime,omitempty"`
	UpdateTime  time.Time `json:"updateTime,omitempty"`
	CreateBy    string    `json:"createBy,omitempty"`
	UpdateBy    string    `json:"updateBy,omitempty"`
	Remark      string    `json:"remark,omitempty"`
}

func (m *SysCreateUserDTO) Convert(user *system.SysUser) {
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
	user.LoginIp = m.LoginIp
	user.LoginDate = m.LoginDate
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}

type SysUpdateUserDTO struct {
	UserId      int64     `json:"userid,omitempty" binding:"required"`
	DeptId      int64     `json:"deptId,omitempty"`
	UserName    string    `json:"username,omitempty"`
	NickName    string    `json:"nickname,omitempty"`
	UserType    string    `json:"userType,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	Sex         string    `json:"sex,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	Password    string    `json:"password,omitempty"`
	Status      string    `json:"status,omitempty"`
	DelFlag     string    `json:"delFlag,omitempty"`
	LoginIp     string    `json:"loginIp,omitempty"`
	LoginDate   time.Time `json:"loginDate,omitempty"`
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
	user.LoginIp = m.LoginIp
	user.LoginDate = m.LoginDate
	user.CreateTime = m.CreateTime
	user.UpdateTime = m.UpdateTime
	user.CreateBy = m.CreateBy
	user.UpdateBy = m.UpdateBy
	user.Remark = m.Remark
}
