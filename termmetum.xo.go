// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// Termmetum represents a row from '` + GetTableName("termmeta") + `'.
type Termmetum struct {
	MetaID    uint64 `json:"meta_id"`    // meta_id
	TermID    uint64 `json:"term_id"`    // term_id
	MetaKey   string `json:"meta_key"`   // meta_key
	MetaValue string `json:"meta_value"` // meta_value

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Termmetum exists in the database.
func (t *Termmetum) Exists() bool {
	return t._exists
}

// Deled provides information if the Termmetum has been deled from the database.
func (t *Termmetum) Deled() bool {
	return t._deleted
}

// Insert inserts the Termmetum to the database.
func (t *Termmetum) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("termmeta") + ` (` +
		`term_id, meta_key, meta_value` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.TermID, t.MetaKey, t.MetaValue)
	res, err := db.Exec(sqlstr, t.TermID, t.MetaKey, t.MetaValue)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	t.MetaID = uint64(id)
	t._exists = true

	return nil
}

// Update updates the Termmetum in the database.
func (t *Termmetum) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("termmeta") + ` SET ` +
		`term_id = ?, meta_key = ?, meta_value = ?` +
		` WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, t.TermID, t.MetaKey, t.MetaValue, t.MetaID)
	_, err = db.Exec(sqlstr, t.TermID, t.MetaKey, t.MetaValue, t.MetaID)
	return err
}

// Save saves the Termmetum to the database.
func (t *Termmetum) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Termmetum from the database.
func (t *Termmetum) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("termmeta") + ` WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, t.MetaID)
	_, err = db.Exec(sqlstr, t.MetaID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// TermmetaByMetaKey retrieves a row from '` + GetTableName("termmeta") + `' as a Termmetum.
//
// Generated from index 'meta_key'.
func TermmetaByMetaKey(db XODB, metaKey string) ([]*Termmetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, term_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("termmeta") + ` ` +
		`WHERE meta_key = ?`

	// run query
	XOLog(sqlstr, metaKey)
	q, err := db.Query(sqlstr, metaKey)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Termmetum{}
	for q.Next() {
		t := Termmetum{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.MetaID, &t.TermID, &t.MetaKey, &t.MetaValue)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// TermmetaByTermID retrieves a row from '` + GetTableName("termmeta") + `' as a Termmetum.
//
// Generated from index 'term_id'.
func TermmetaByTermID(db XODB, termID uint64) ([]*Termmetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, term_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("termmeta") + ` ` +
		`WHERE term_id = ?`

	// run query
	XOLog(sqlstr, termID)
	q, err := db.Query(sqlstr, termID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Termmetum{}
	for q.Next() {
		t := Termmetum{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.MetaID, &t.TermID, &t.MetaKey, &t.MetaValue)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// TermmetumByMetaID retrieves a row from '` + GetTableName("termmeta") + `' as a Termmetum.
//
// Generated from index 'termmeta_meta_id_pkey'.
func TermmetumByMetaID(db XODB, metaID uint64) (*Termmetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, term_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("termmeta") + ` ` +
		`WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, metaID)
	t := Termmetum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, metaID).Scan(&t.MetaID, &t.TermID, &t.MetaKey, &t.MetaValue)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
