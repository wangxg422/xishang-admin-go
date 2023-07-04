package system

import "time"

type SysUser struct {
	UserId      int64     `gorm:"primaryKey;column:user_id" json:"userId,omitempty"`
	DeptId      int64     `gorm:"column:dept_id" json:"deptId,omitempty"`
	UserName    string    `gorm:"column:user_name" json:"userName,omitempty"`
	NickName    string    `gorm:"column:nick_name" json:"nickName,omitempty"`
	UserType    string    `gorm:"column:user_type" json:"userType,omitempty"`
	Email       string    `gorm:"column:email" json:"email,omitempty"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber,omitempty"`
	Sex         int8      `gorm:"column:sex" json:"sex"`
	Avatar      string    `gorm:"column:avatar" json:"avatar,omitempty"` //头像地址
	Password    string    `gorm:"column:password" json:"-"`
	Status      int8      `gorm:"column:status;default:0" json:"status"`
	DelFlag     int8      `gorm:"column:del_flag;default:0" json:"delFlag"`
	CreateTime  time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime  time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy    string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy    string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark      string    `gorm:"column:remark" json:"remark,omitempty"`

	SysDept  SysDept   `gorm:"foreignKey:DeptId;references:DeptId;comment:用户归属部门" json:"dept,omitempty"`
	SysRoles []SysRole `gorm:"many2many:sys_user_role;foreignKey:UserId;joinForeignKey:user_id;references:RoleId;joinReferences:role_id;" json:"roles,omitempty"`
	SysPosts []SysPost `gorm:"many2many:sys_user_post;foreignKey:UserId;joinForeignKey:user_id;references:PostId;joinReferences:post_id;" json:"posts,omitempty"`
}
