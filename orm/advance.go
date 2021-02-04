package orm

import "github.com/hashicorp/go-multierror"

func (t *MaDB) Exec(sql string, values ...interface{}) *MaDB {
	clone := t.db.Exec(sql, values)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t
}

func (t *MaDB) Preload(query string, args ...interface{}) *MaDB {
	clone := t.db.Preload(query, args)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t
}
