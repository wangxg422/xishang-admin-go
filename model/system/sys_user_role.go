package system

type SysUserRole struct {
	UserId int64 `gorm:"column:user_id;"`
	RoleId int64 `gorm:"column:role_id;"`
}
