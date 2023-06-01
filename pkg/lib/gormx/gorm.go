package gormx

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var (
	JudgeryDB *gorm.DB
)

var (
	db  *gorm.DB
	err error
)

func ConnectWithDSN(dsn string, conf DatabaseConf) (*gorm.DB, error) {
	if strings.HasPrefix(dsn, "mysql://") {
		db, err = gorm.Open(mysql.Open(strings.ReplaceAll(dsn, "mysql://", "")))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, gorm.ErrUnsupportedDriver
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if conf.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
	}
	if conf.MaxIdleConn > 0 {
		sqlDB.SetMaxOpenConns(conf.MaxOpenConn)
	}
	if conf.ConnMaxLiftTime > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLiftTime) * time.Second)
	}
	return db, nil

}

//ConnectJudgeryDB 根据conf创建全局DB
func ConnectJudgeryDB(conf DatabaseConf) error {
	db, err := ConnectWithDSN(conf.JudgeryDSN, conf)
	if err != nil {
		return err
	}
	//JudgeryDB = db.Debug()	//开启sql debug模式,会终端打印sql执行日志
	JudgeryDB = db
	return nil
}
