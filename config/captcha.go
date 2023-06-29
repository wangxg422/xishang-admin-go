package config

type Captcha struct {
	Enabled    bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	KeyLong    int    `mapstructure:"key-long" json:"key-long" yaml:"key-long"`
	ImgWidth   int    `mapstructure:"img-width" json:"img-width" yaml:"img-width"`
	ImgHeight  int    `mapstructure:"img-height" json:"img-height" yaml:"img-height"`
	Expiration int64  `mapstructure:"expiration" json:"expiration" yaml:"expiration"`
	PreKey     string `mapstructure:"pre-key" json:"pre-key" yaml:"pre-key"`
}
