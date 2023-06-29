package system

type JwtBlacklist struct {
	Jwt string `gorm:"type:text;comment:jwt"`
}
