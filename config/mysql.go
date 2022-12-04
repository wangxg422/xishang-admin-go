package config

type Mysql struct {
	Address     string `mapstructure:"address" json:"address" yaml:"address"`                   // 服务器地址:端口
	Port        string `mapstructure:"port" json:"port" yaml:"port"`                            //:端口
	ConnConfig  string `mapstructure:"conn-config" json:"conn-config" yaml:"conn-config"`       // 高级配置
	Dbname      string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                   // 数据库名
	Username    string `mapstructure:"username" json:"username" yaml:"username"`                // 数据库用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`                // 数据库密码
	TablePrefix string `mapstructure:"table-prefix" json:"table-prefix" yaml:"table-prefix"`    //全局表前缀，单独定义TableName则不生效
	Singular    bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                //是否开启全局禁用复数，true表示开启
	Engine      string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`     //数据库引擎，默认InnoDB
	MaxIdleConn int    `mapstructure:"max-idle-conn" json:"max-idle-conn" yaml:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"max-open-conn" json:"max-open-conn" yaml:"max-open-conn"` // 打开到数据库的最大连接数
	LogMode     string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                // 是否开启Gorm全局日志
	LogZap      bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                   // 是否通过zap写入日志文件
}
