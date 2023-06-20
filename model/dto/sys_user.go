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

func (u *SysCreateUserDTO) Convert(user *system.SysUser) {
	user.DeptId = u.DeptId
	user.UserName = u.UserName
	user.NickName = u.NickName
	user.UserType = u.UserType
	user.Email = u.Email
	user.PhoneNumber = u.PhoneNumber
	user.Sex = u.Sex
	user.Avatar = u.Avatar
	user.Password = u.Password
	user.Status = u.Status
	user.DelFlag = u.DelFlag
	user.LoginIp = u.LoginIp
	user.LoginDate = u.LoginDate
	user.CreateTime = u.CreateTime
	user.UpdateTime = u.UpdateTime
	user.CreateBy = u.CreateBy
	user.UpdateBy = u.UpdateBy
	user.Remark = u.Remark
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

func (u *SysUpdateUserDTO) Convert(user *system.SysUser) {
	user.UserId = u.UserId
	user.DeptId = u.DeptId
	user.UserName = u.UserName
	user.NickName = u.NickName
	user.UserType = u.UserType
	user.Email = u.Email
	user.PhoneNumber = u.PhoneNumber
	user.Sex = u.Sex
	user.Avatar = u.Avatar
	user.Password = u.Password
	user.Status = u.Status
	user.DelFlag = u.DelFlag
	user.LoginIp = u.LoginIp
	user.LoginDate = u.LoginDate
	user.CreateTime = u.CreateTime
	user.UpdateTime = u.UpdateTime
	user.CreateBy = u.CreateBy
	user.UpdateBy = u.UpdateBy
	user.Remark = u.Remark
}
