package config

type AppConfig struct {
	App     App     `mapstructure:"app" json:"app" yaml:"app"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
