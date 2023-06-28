package config

type Redis struct {
	Addresses string `mapstructure:"address" json:"address" yaml:"address"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	Password  string `mapstructure:"password" json:"password" yaml:"password"`
	Db        int    `mapstructure:"db" json:"db" yaml:"db"`
}
