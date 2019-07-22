// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/patrickmn/go-cache"

	"github.com/95key/util"
)

// TermRelationship represents a row from '` + GetTableName("term_relationships") + `'.
type TermRelationship struct {
	ObjectID       uint64 `json:"object_id"`        // object_id
	TermTaxonomyID uint64 `json:"term_taxonomy_id"` // term_taxonomy_id
	TermOrder      int    `json:"term_order"`       // term_order

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TermRelationship exists in the database.
func (tr *TermRelationship) Exists() bool {
	return tr._exists
}

// Deled provides information if the TermRelationship has been deled from the database.
func (tr *TermRelationship) Deled() bool {
	return tr._deleted
}

// Insert inserts the TermRelationship to the database.
func (tr *TermRelationship) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if tr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	var sqlstr = `INSERT INTO ` + GetTableName("term_relationships") + ` (` +
		`object_id, term_taxonomy_id, term_order` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, tr.ObjectID, tr.TermTaxonomyID, tr.TermOrder)
	_, err = db.Exec(sqlstr, tr.ObjectID, tr.TermTaxonomyID, tr.TermOrder)
	if err != nil {
		return err
	}

	// set existence
	tr._exists = true

	return nil
}

// Update updates the TermRelationship in the database.
func (tr *TermRelationship) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if tr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query with composite primary key
	var sqlstr = `UPDATE ` + GetTableName("term_relationships") + ` SET ` +
		`term_order = ?` +
		` WHERE object_id = ? AND term_taxonomy_id = ?`

	// run query
	XOLog(sqlstr, tr.TermOrder, tr.ObjectID, tr.TermTaxonomyID)
	_, err = db.Exec(sqlstr, tr.TermOrder, tr.ObjectID, tr.TermTaxonomyID)
	return err
}

// Save saves the TermRelationship to the database.
func (tr *TermRelationship) Save(db XODB) error {
	if tr.Exists() {
		return tr.Update(db)
	}

	return tr.Insert(db)
}

// Delete deletes the TermRelationship from the database.
func (tr *TermRelationship) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tr._exists {
		return nil
	}

	// if deleted, bail
	if tr._deleted {
		return nil
	}

	// sql query with composite primary key
	var sqlstr = `DELETE FROM ` + GetTableName("term_relationships") + ` WHERE object_id = ? AND term_taxonomy_id = ?`

	// run query
	XOLog(sqlstr, tr.ObjectID, tr.TermTaxonomyID)
	_, err = db.Exec(sqlstr, tr.ObjectID, tr.TermTaxonomyID)
	if err != nil {
		return err
	}

	// set deleted
	tr._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// TermRelationshipByTermTaxonomyID retrieves a row from '` + GetTableName("term_relationships") + `' as a TermRelationship.
//
// Generated from index 'term_relationships_term_taxonomy_id_pkey'.
func TermRelationshipByTermTaxonomyID(db XODB, termTaxonomyID uint64) (*TermRelationship, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`object_id, term_taxonomy_id, term_order ` +
		`FROM ` + GetTableName("term_relationships") + ` ` +
		`WHERE term_taxonomy_id = ?`

	// run query
	XOLog(sqlstr, termTaxonomyID)
	tr := TermRelationship{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, termTaxonomyID).Scan(&tr.ObjectID, &tr.TermTaxonomyID, &tr.TermOrder)
	if err != nil {
		return nil, err
	}

	return &tr, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// TermRelationshipsByTermTaxonomyID retrieves a row from '` + GetTableName("term_relationships") + `' as a TermRelationship.
//
// Generated from index 'term_taxonomy_id'.
func TermRelationshipsByTermTaxonomyID(db XODB, termTaxonomyID uint64) ([]*TermRelationship, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`object_id, term_taxonomy_id, term_order ` +
		`FROM ` + GetTableName("term_relationships") + ` ` +
		`WHERE term_taxonomy_id = ?`

	// run query
	XOLog(sqlstr, termTaxonomyID)
	q, err := db.Query(sqlstr, termTaxonomyID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TermRelationship{}
	for q.Next() {
		tr := TermRelationship{
			_exists: true,
		}

		// scan
		err = q.Scan(&tr.ObjectID, &tr.TermTaxonomyID, &tr.TermOrder)
		if err != nil {
			return nil, err
		}

		res = append(res, &tr)
	}

	return res, nil
}

// TermRelationshipsByObjectID { {- $table := (schema .Schema .Type.Table.TableName) -} }// TermRelationshipsByTermTaxonomyID retrieves a row from '` + GetTableName("term_relationships") + `' as a TermRelationship.
//
// Generated from index 'term_taxonomy_id'.
// func TermRelationshipsByObjectID(db wpsapi.XODB, objectID uint64) ([]*TermRelationship, error) {
func TermRelationshipsByObjectID(db XODB, objectID uint64) ([]*TermRelationship, error) {
	v, found := c.Get("/term_relationships/object_id/" + util.Uint64ToString(objectID))
	if found {
		return v.([]*TermRelationship), nil
	}
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`object_id, term_taxonomy_id, term_order ` +
		`FROM ` + GetTableName("term_relationships") + ` ` +
		`WHERE object_id = ?`

	// run query
	XOLog(sqlstr, objectID)
	q, err := db.Query(sqlstr, objectID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := make([]*TermRelationship, 0)
	for q.Next() {

		tr := TermRelationship{}

		// scan
		err = q.Scan(&tr.ObjectID, &tr.TermTaxonomyID, &tr.TermOrder)
		if err != nil {
			return nil, err
		}

		res = append(res, &tr)
	}

	c.Set("/term_relationships/object_id/"+util.Uint64ToString(objectID), res, cache.DefaultExpiration)

	return res, nil
}
