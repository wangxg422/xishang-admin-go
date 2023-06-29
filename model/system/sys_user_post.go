package system

type SysUserPost struct {
	UserId int64 `gorm:"column:user_id"`
	PostId int64 `gorm:"column:post_id"`
}
