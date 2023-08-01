package system

import "time"

type SysUser struct {
	UserId      int64     `gorm:"primaryKey;column:user_id" json:"userId"`
	DeptId      int64     `gorm:"column:dept_id" json:"deptId"`
	UserName    string    `gorm:"column:user_name" json:"userName"`
	UserNumber  string    `gorm:"column:user_number" json:"userNumber"`
	RealName    string    `gorm:"column:real_name" json:"realName"`
	NickName    string    `gorm:"column:nick_name" json:"nickName"`
	UserType    string    `gorm:"column:user_type" json:"userType"`
	Email       string    `gorm:"column:email" json:"email"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	Sex         string    `gorm:"column:sex" json:"sex"`
	Avatar      string    `gorm:"column:avatar" json:"avatar"`
	Password    string    `gorm:"column:password" json:"-"`
	Status      string    `gorm:"column:status" json:"status"`
	DelFlag     int8      `gorm:"column:del_flag" json:"delFlag"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreateBy    string    `gorm:"column:create_by" json:"createBy"`
	UpdateBy    string    `gorm:"column:update_by" json:"updateBy"`
	Remark      string    `gorm:"column:remark" json:"remark"`

	SysDept  SysDept   `gorm:"foreignKey:DeptId;references:DeptId" json:"dept,omitempty"`
	SysRoles []SysRole `gorm:"many2many:sys_user_role;foreignKey:UserId;joinForeignKey:UserId;references:RoleId;joinReferences:RoleId;" json:"roles,omitempty"`
	SysPosts []SysPost `gorm:"many2many:sys_user_post;foreignKey:UserId;joinForeignKey:UserId;references:PostId;joinReferences:PostId;" json:"posts,omitempty"`
}
