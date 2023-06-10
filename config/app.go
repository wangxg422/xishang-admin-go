package config

type App struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
