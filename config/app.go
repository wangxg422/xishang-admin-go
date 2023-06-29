package config

type App struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`
	Addresses     string `mapstructure:"addresses" json:"addresses" yaml:"addresses"`
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
}
