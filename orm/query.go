package orm

import (
	"database/sql"
	"github.com/hashicorp/go-multierror"
	"gorm.io/gorm"
)

// db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
// db.Find(&users, User{Age: 20})
//// SELECT * FROM users WHERE age = 20;

// Map
// db.Find(&users, map[string]interface{}{"age": 20})
//// SELECT * FROM users WHERE age = 20;

// db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// db.Where("role = ?", "admin").Or("role = ?", "super_admin").Not("name = ?", "jinzhu").Find(&users)
func (t *MaDB) Find(entitiesPrt interface{}, where ...interface{}) *MaDB {
	clone := t.db.Find(entitiesPrt, where...)

	if clone.Error != nil && clone.Error != gorm.ErrRecordNotFound {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.RowsAffected = clone.RowsAffected

	t.db = clone

	return t
}

// Scan scan value to a struct
// var result Result
// db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result)

// Raw SQL
// db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
func (t *MaDB) Scan(dest interface{}) *MaDB {
	clone := t.db.Scan(dest)

	if clone.Error != nil && clone.Error != gorm.ErrRecordNotFound {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.RowsAffected = clone.RowsAffected

	t.db = clone

	return t
}

// Pluck used to query single column from a model as a map
// 将模型中的单个列作为地图查询，如果要查询多个列，可以使用Scan
// var ages []int64
// db.Find(&users).Pluck("age", &ages)

// var names []string
// db.Model(&User{}).Pluck("name", &names)

// db.Table("deleted_users").Pluck("name", &names)

// 要返回多个列，做这样：
//db.Select("name, age").Find(&users)
func (t *MaDB) Pluck(column string, value interface{}) *MaDB {
	clone := t.db.Pluck(column, value)

	if clone.Error != nil && clone.Error != gorm.ErrRecordNotFound {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.RowsAffected = clone.RowsAffected

	t.db = clone

	return t
}

// 查找第一行数据，按照主键排序 asc
// SELECT * FROM `table` ORDER BY `id` ASC LIMIT 1
// var item model.OrderFilterData
// qa.First(&item)
// SELECT * FROM `mpv_order_filter_data`  WHERE (`id` = 100) ORDER BY `id` ASC LIMIT 1
// qa.First(&item,100)
// SELECT * FROM `mpv_order_filter_data`  WHERE (name = 'xxx') ORDER BY `id` ASC LIMIT 1
// qa.First(&item,"name = ?","xxx")
func (t *MaDB) First(entityPrt interface{}, where ...interface{}) *MaDB {
	clone := t.db.First(entityPrt, where...)

	if clone.Error != nil && clone.Error != gorm.ErrRecordNotFound {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.RowsAffected = clone.RowsAffected

	t.db = clone

	return t
}

// 随机获取一条记录
// db.Take(&user)
// SELECT * FROM users LIMIT 1;
func (t *MaDB) Take(out interface{}, where ...interface{}) *MaDB {
	clone := t.db.Take(out, where...)

	if clone.Error != nil && clone.Error != gorm.ErrRecordNotFound {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.RowsAffected = clone.RowsAffected

	t.db = clone

	return t
}

// Raw SQL
// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows() // (*sql.Rows, error)
// defer rows.Close()
// for rows.Next() {
// rows.Scan(&name, &age, &email)
// }
func (t *MaDB) Rows() (*sql.Rows, error) {
	return t.db.Rows()
}

// row := db.Table("users").Where("name = ?", "jinzhu").Select("name, age").Row() // (*sql.Row)
// row.Scan(&name, &age)
func (t *MaDB) Row() *sql.Row {
	return t.db.Row()
}
