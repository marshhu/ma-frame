package orm

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func Init(conf *DbSettings) error {
	defaultKey = conf.DbName
	return InitDBWithDbName(conf, conf.DbName)
}

func InitDBWithDbName(conf *DbSettings, dbName string) error {
	// 初始化数据库可以在程序中动态调用，需要防止并发
	chk := CheckDbExists(dbName)
	if chk == true {
		return nil
	}
	maxIdel := 5
	if conf.MaxIdleConns > 0 {
		maxIdel = conf.MaxIdleConns
	}
	maxConn := 20
	if conf.MaxOpenConns > 0 {
		maxConn = conf.MaxOpenConns
	}
	lifeTime := 5
	if conf.ConnMaxLifetime > 0 {
		lifeTime = conf.ConnMaxLifetime
	}

	p := strconv.Itoa(conf.Port)
	var dialector gorm.Dialector
	switch conf.Dialect {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=true&loc=Local", conf.User, conf.Password, conf.Host, p, conf.DbName)
		dialector = mysql.Open(dsn)
	case "mssql":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=600&encrypt=disable", conf.User, conf.Password, conf.Host, p, conf.DbName)
		dialector = sqlserver.Open(dsn)
	default:
		return fmt.Errorf("不支持的数据库类型：%s", conf.Dialect)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return errors.Wrapf(err, "无法打开数据库连接")
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdel)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxConn)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(lifeTime))

	// 保存DB实例
	setDbInstance(dbName, db)

	return nil
}
