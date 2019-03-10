// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Sitemetum represents a row from '` + GetTableName("sitemeta") + `'.
type Sitemetum struct {
	MetaID    int64          `json:"meta_id"`    // meta_id
	SiteID    int64          `json:"site_id"`    // site_id
	MetaKey   sql.NullString `json:"meta_key"`   // meta_key
	MetaValue sql.NullString `json:"meta_value"` // meta_value

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Sitemetum exists in the database.
func (s *Sitemetum) Exists() bool {
	return s._exists
}

// Deled provides information if the Sitemetum has been deled from the database.
func (s *Sitemetum) Deled() bool {
	return s._deleted
}

// Insert inserts the Sitemetum to the database.
func (s *Sitemetum) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("sitemeta") + ` (` +
		`site_id, meta_key, meta_value` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, s.SiteID, s.MetaKey, s.MetaValue)
	res, err := db.Exec(sqlstr, s.SiteID, s.MetaKey, s.MetaValue)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	s.MetaID = int64(id)
	s._exists = true

	return nil
}

// Update updates the Sitemetum in the database.
func (s *Sitemetum) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if s._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("sitemeta") + ` SET ` +
		`site_id = ?, meta_key = ?, meta_value = ?` +
		` WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, s.SiteID, s.MetaKey, s.MetaValue, s.MetaID)
	_, err = db.Exec(sqlstr, s.SiteID, s.MetaKey, s.MetaValue, s.MetaID)
	return err
}

// Save saves the Sitemetum to the database.
func (s *Sitemetum) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Delete deletes the Sitemetum from the database.
func (s *Sitemetum) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !s._exists {
		return nil
	}

	// if deleted, bail
	if s._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("sitemeta") + ` WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, s.MetaID)
	_, err = db.Exec(sqlstr, s.MetaID)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SitemetaByMetaKey retrieves a row from '` + GetTableName("sitemeta") + `' as a Sitemetum.
//
// Generated from index 'meta_key'.
func SitemetaByMetaKey(db XODB, metaKey sql.NullString) ([]*Sitemetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, site_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("sitemeta") + ` ` +
		`WHERE meta_key = ?`

	// run query
	XOLog(sqlstr, metaKey)
	q, err := db.Query(sqlstr, metaKey)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Sitemetum{}
	for q.Next() {
		s := Sitemetum{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.MetaID, &s.SiteID, &s.MetaKey, &s.MetaValue)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SitemetaBySiteID retrieves a row from '` + GetTableName("sitemeta") + `' as a Sitemetum.
//
// Generated from index 'site_id'.
func SitemetaBySiteID(db XODB, siteID int64) ([]*Sitemetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, site_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("sitemeta") + ` ` +
		`WHERE site_id = ?`

	// run query
	XOLog(sqlstr, siteID)
	q, err := db.Query(sqlstr, siteID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Sitemetum{}
	for q.Next() {
		s := Sitemetum{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.MetaID, &s.SiteID, &s.MetaKey, &s.MetaValue)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SitemetumByMetaID retrieves a row from '` + GetTableName("sitemeta") + `' as a Sitemetum.
//
// Generated from index 'sitemeta_meta_id_pkey'.
func SitemetumByMetaID(db XODB, metaID int64) (*Sitemetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, site_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("sitemeta") + ` ` +
		`WHERE meta_id = ?`

	// run query
	XOLog(sqlstr, metaID)
	s := Sitemetum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, metaID).Scan(&s.MetaID, &s.SiteID, &s.MetaKey, &s.MetaValue)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
