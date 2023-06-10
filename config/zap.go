package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`    // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"` // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"` // 输出
	Path          string `mapstructure:"path" json:"path"  yaml:"path"`
	FileShunt     bool   `mapstructure:"file-shunt" json:"file-shunt"  yaml:"file-shunt"`            // 日志文件夹
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名

	MaxAge       int  `mapstructure:"max-age" json:"max-age" yaml:"max-age"`
	MaxBackups   int  `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`
	MaxSize      int  `mapstructure:"max-size" json:"max-size" yaml:"max-size"`                   // 日志留存时间
	LogInConsole bool `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
}
