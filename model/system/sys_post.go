package system

import "time"

type SysPost struct {
	PostId     int64     `gorm:"primaryKey;column:post_id" json:"postId,omitempty"`
	PostCode   string    `gorm:"column:post_code" json:"postCode,omitempty"`
	PostName   string    `gorm:"column:post_name" json:"postName,omitempty"`
	PostSort   int       `gorm:"column:post_sort" json:"postSort,omitempty"`
	Status     string    `gorm:"column:status;default:0" json:"status,omitempty"`
	DelFlag    int8      `gorm:"column:del_flag;default:0" json:"delFlag,omitempty"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime,omitempty"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime,omitempty"`
	CreateBy   string    `gorm:"column:create_by" json:"createBy,omitempty"`
	UpdateBy   string    `gorm:"column:update_by" json:"updateBy,omitempty"`
	Remark     string    `gorm:"column:remark" json:"remark,omitempty"`

	SysUsers []SysUser `gorm:"many2many:sys_user_post;foreignKey:PostId;joinForeignKey:PostId;references:UserId;joinReferences:UserId;" json:"users,omitempty"`
}
