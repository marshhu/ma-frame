package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"sync"
)

// DbSetting 数据库连接字符串属性
type DbSettings struct {
	Dialect         string
	DbName          string
	Host            string
	User            string
	Password        string
	Port            int
	Key             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // 分钟
}

type MaDB struct{
	db *gorm.DB
	dbName           string
	Error            error
	RowsAffected     int64 // 受影响行数
}

var defaultKey string

// 连接实例池
// var dbInstances map[string]*gorm.DB
var dbInstances = struct {
	sync.RWMutex
	m map[string]*gorm.DB
}{m: make(map[string]*gorm.DB)}

// DB 返回一个连接池的实例
// var db *gorm.DB

func getDbInstance(key string) *gorm.DB {
	dbInstances.RLock()
	defer dbInstances.RUnlock()

	db := dbInstances.m[key]
	return db
}

func CheckDbExists(key string) bool {
	dbInstances.RLock()
	defer dbInstances.RUnlock()

	_, ok := dbInstances.m[key]
	return ok
}

func defaultDb() *gorm.DB {
	dbInstances.RLock()
	defer dbInstances.RUnlock()

	dbInst, ok := dbInstances.m[defaultKey]

	if ok == true {
		return dbInst
	} else {
		panic("没有发现DB实例" + defaultKey)
	}
}

func setDbInstance(key string, db *gorm.DB) {
	dbInstances.Lock()
	defer dbInstances.Unlock()

	if _, ok := dbInstances.m[key]; ok == false {
		dbInstances.m[key] = db
	}
}

// 移除DB实例
func RemoveDbByKey(key string) {
	dbInstances.RLock()
	defer dbInstances.RUnlock()

	_, ok := dbInstances.m[key]
	if ok == true {
		delete(dbInstances.m, key)
	}
}

// NewQuery 获取一个查询会话（只作用于一次查询）
func NewQuery() *MaDB {
	var maDb = new(MaDB)
	maDb.db = defaultDb()
	maDb.dbName = defaultKey
	return maDb
}