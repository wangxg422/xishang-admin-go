package global

import (
	"backend/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Viper     *viper.Viper
	AppConfig config.AppConfig
)
