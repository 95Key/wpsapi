// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// BlogVersion represents a row from '` + GetTableName("blog_versions") + `'.
type BlogVersion struct {
	BlogID      int64     `json:"blog_id"`      // blog_id
	DbVersion   string    `json:"db_version"`   // db_version
	LastUpdated time.Time `json:"last_updated"` // last_updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BlogVersion exists in the database.
func (bv *BlogVersion) Exists() bool {
	return bv._exists
}

// Deled provides information if the BlogVersion has been deled from the database.
func (bv *BlogVersion) Deled() bool {
	return bv._deleted
}

// Insert inserts the BlogVersion to the database.
func (bv *BlogVersion) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if bv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	var sqlstr = `INSERT INTO ` + GetTableName("blog_versions") + ` (` +
		`blog_id, db_version, last_updated` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, bv.BlogID, bv.DbVersion, bv.LastUpdated)
	_, err = db.Exec(sqlstr, bv.BlogID, bv.DbVersion, bv.LastUpdated)
	if err != nil {
		return err
	}

	// set existence
	bv._exists = true

	return nil
}

// Update updates the BlogVersion in the database.
func (bv *BlogVersion) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if bv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("blog_versions") + ` SET ` +
		`db_version = ?, last_updated = ?` +
		` WHERE blog_id = ?`

	// run query
	XOLog(sqlstr, bv.DbVersion, bv.LastUpdated, bv.BlogID)
	_, err = db.Exec(sqlstr, bv.DbVersion, bv.LastUpdated, bv.BlogID)
	return err
}

// Save saves the BlogVersion to the database.
func (bv *BlogVersion) Save(db XODB) error {
	if bv.Exists() {
		return bv.Update(db)
	}

	return bv.Insert(db)
}

// Delete deletes the BlogVersion from the database.
func (bv *BlogVersion) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bv._exists {
		return nil
	}

	// if deleted, bail
	if bv._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("blog_versions") + ` WHERE blog_id = ?`

	// run query
	XOLog(sqlstr, bv.BlogID)
	_, err = db.Exec(sqlstr, bv.BlogID)
	if err != nil {
		return err
	}

	// set deleted
	bv._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// BlogVersionByBlogID retrieves a row from '` + GetTableName("blog_versions") + `' as a BlogVersion.
//
// Generated from index 'blog_versions_blog_id_pkey'.
func BlogVersionByBlogID(db XODB, blogID int64) (*BlogVersion, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`blog_id, db_version, last_updated ` +
		`FROM ` + GetTableName("blog_versions") + ` ` +
		`WHERE blog_id = ?`

	// run query
	XOLog(sqlstr, blogID)
	bv := BlogVersion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, blogID).Scan(&bv.BlogID, &bv.DbVersion, &bv.LastUpdated)
	if err != nil {
		return nil, err
	}

	return &bv, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// BlogVersionsByDbVersion retrieves a row from '` + GetTableName("blog_versions") + `' as a BlogVersion.
//
// Generated from index 'db_version'.
func BlogVersionsByDbVersion(db XODB, dbVersion string) ([]*BlogVersion, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`blog_id, db_version, last_updated ` +
		`FROM ` + GetTableName("blog_versions") + ` ` +
		`WHERE db_version = ?`

	// run query
	XOLog(sqlstr, dbVersion)
	q, err := db.Query(sqlstr, dbVersion)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*BlogVersion{}
	for q.Next() {
		bv := BlogVersion{
			_exists: true,
		}

		// scan
		err = q.Scan(&bv.BlogID, &bv.DbVersion, &bv.LastUpdated)
		if err != nil {
			return nil, err
		}

		res = append(res, &bv)
	}

	return res, nil
}