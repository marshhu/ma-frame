package orm

import "github.com/hashicorp/go-multierror"

func (t *MaDB) Create(entityPtr interface{}) error {
	clone := t.db.Create(entityPtr)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}

func (t *MaDB) Update(column string, value interface{}) error {
	clone := t.db.Update(column, value)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}

func (t *MaDB) UpdateColumn(column string, value interface{}) error {
	clone := t.db.UpdateColumn(column, value)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}

func (t *MaDB) Updates(values interface{}) error {
	clone := t.db.Updates(values)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}

func (t *MaDB) UpdateColumns(values interface{}) error {
	clone := t.db.UpdateColumns(values)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}

func (t *MaDB) Delete(value interface{}, conds ...interface{}) error {
	clone := t.db.Delete(value, conds)
	if clone.Error != nil {
		t.Error = multierror.Append(t.Error, clone.Error)
	}

	t.db = clone
	t.RowsAffected = clone.RowsAffected
	return t.Error
}
