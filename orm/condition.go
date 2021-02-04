package orm

import multierror "github.com/hashicorp/go-multierror"

// Table 设置表名
func (t *MaDB) Table(tableName string) *MaDB {
	clone := t.db.Table(tableName)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}
	t.db = clone
	return t
}

// Select 返回column字段
// db.Select("name, age").Find(&users)
//// SELECT name, age FROM users;

// db.Select([]string{"name", "age"}).Find(&users)
//// SELECT name, age FROM users;

// db.Table("users").Select("COALESCE(age,?)", 42).Rows()
//// SELECT COALESCE(age,'42') FROM users;
func (t *MaDB) Select(query interface{}, args ...interface{}) *MaDB {
	clone := t.db.Select(query, args...)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// var result Result
// db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
func (t *MaDB) Raw(sql string, values ...interface{}) *MaDB {
	clone := t.db.Raw(sql, values...)

	t.RowsAffected = clone.RowsAffected

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}


// Joins specify Joins conditions
// rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
// db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
func (t *MaDB) Joins(query string, args ...interface{}) *MaDB {
	clone := t.db.Joins(query, args...)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// Group specify the group method on the find
// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
// db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
func (t *MaDB) Group(query string) *MaDB {
	clone := t.db.Group(query)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// Having specify HAVING conditions for GROUP BY
// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
// db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)
func (t *MaDB) Having(query interface{}, values ...interface{}) *MaDB {
	clone := t.db.Having(query, values...)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// 获取模型的记录数
// db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

// db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

// db.Table("deleted_users").Count(&count)
//// SELECT count(*) FROM deleted_users;
func (t *MaDB) Count(value *int64) *MaDB {
	clone := t.db.Count(value)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// db.Limit(3).Find(&users)
//// SELECT * FROM users LIMIT 3;

// Cancel limit condition with -1
// db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
//// SELECT * FROM users LIMIT 10; (users1)
//// SELECT * FROM users; (users2)
func (t *MaDB) Limit(limit int) *MaDB {
	clone := t.db.Limit(limit)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// 指定在开始返回记录之前要跳过的记录数
// db.Offset(3).Find(&users)
//// SELECT * FROM users OFFSET 3;

// Cancel offset condition with -1
// db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
//// SELECT * FROM users OFFSET 10; (users1)
//// SELECT * FROM users; (users2)
func (t *MaDB) Offset(offset int) *MaDB {
	clone := t.db.Offset(offset)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// db.Order("age desc, name").Find(&users)
//// SELECT * FROM users ORDER BY age desc, name;

// Multiple orders
// db.Order("age desc").Order("name").Find(&users)
//// SELECT * FROM users ORDER BY age desc, name;

// ReOrder
// db.Order("age desc").Find(&users1)
// SELECT * FROM users ORDER BY age desc; (users1)
func (t *MaDB) Order(value interface{}) *MaDB {
	clone := t.db.Order(value)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
// db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

// Map
// db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
func (t *MaDB) Or(query interface{}, args ...interface{}) *MaDB {
	clone := t.db.Or(query, args...)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// db.Not("name", "jinzhu").First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

// Not In
// db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

// Not In slice of primary keys
// db.Not([]int64{1,2,3}).First(&user)
//// SELECT * FROM users WHERE id NOT IN (1,2,3);

// Plain SQL
// db.Not("name = ?", "jinzhu").First(&user)
//// SELECT * FROM users WHERE NOT(name = "jinzhu");

// Struct
// db.Not(User{Name: "jinzhu"}).First(&user)
//// SELECT * FROM users WHERE name <> "jinzhu";
func (t *MaDB) Not(query interface{}, args ...interface{}) *MaDB {
	clone := t.db.Not(query, args...)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}

// Model specify the model you would like to run db operations
//    // update all users's name to `hello`
//    db.Model(&User{}).Update("name", "hello")
//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
//    db.Model(&user).Update("name", "hello")
func (t *MaDB) Model(entityPtr interface{}) *MaDB {
	clone := t.db.Model(entityPtr)

	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone

	return t
}