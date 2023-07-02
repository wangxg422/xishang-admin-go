package initial

import (
	"backend/global"
	"fmt"
	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	dbConfig := global.AppConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", dbConfig.Username, dbConfig.Password, dbConfig.Address, dbConfig.Port, dbConfig.Dbname, dbConfig.ConnConfig)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), gormConfig())

	if err != nil {
		panic("init gorm ")
	}
	sqlDb, _ := db.DB()

	sqlDb.SetMaxOpenConns(global.AppConfig.Mysql.MaxOpenConn)
	sqlDb.SetMaxIdleConns(global.AppConfig.Mysql.MaxIdleConn)

	return db
}

func gormConfig() *gorm.Config {
	dbConfig := global.AppConfig.Mysql

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbConfig.TablePrefix,
			SingularTable: dbConfig.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
	return gormConfig
}
