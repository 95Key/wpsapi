// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/patrickmn/go-cache"
	// "util"
)

// Option represents a row from '` + GetTableName("options") + `'.
type Option struct {
	OptionID    uint64 `json:"option_id"`    // option_id
	OptionName  string `json:"option_name"`  // option_name
	OptionValue string `json:"option_value"` // option_value
	Autoload    string `json:"autoload"`     // autoload

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Option exists in the database.
func (o *Option) Exists() bool {
	return o._exists
}

// Deled provides information if the Option has been deled from the database.
func (o *Option) Deled() bool {
	return o._deleted
}

// Insert inserts the Option to the database.
func (o *Option) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if o._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("options") + ` (` +
		`option_name, option_value, autoload` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, o.OptionName, o.OptionValue, o.Autoload)
	res, err := db.Exec(sqlstr, o.OptionName, o.OptionValue, o.Autoload)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	o.OptionID = uint64(id)
	o._exists = true

	return nil
}

// Update updates the Option in the database.
func (o *Option) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if o._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("options") + ` SET ` +
		`option_name = ?, option_value = ?, autoload = ?` +
		` WHERE option_id = ?`

	// run query
	XOLog(sqlstr, o.OptionName, o.OptionValue, o.Autoload, o.OptionID)
	_, err = db.Exec(sqlstr, o.OptionName, o.OptionValue, o.Autoload, o.OptionID)
	return err
}

// Save saves the Option to the database.
func (o *Option) Save(db XODB) error {
	if o.Exists() {
		return o.Update(db)
	}

	return o.Insert(db)
}

// Delete deletes the Option from the database.
func (o *Option) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !o._exists {
		return nil
	}

	// if deleted, bail
	if o._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("options") + ` WHERE option_id = ?`

	// run query
	XOLog(sqlstr, o.OptionID)
	_, err = db.Exec(sqlstr, o.OptionID)
	if err != nil {
		return err
	}

	// set deleted
	o._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// OptionByOptionName retrieves a row from '` + GetTableName("options") + `' as a Option.
//
// Generated from index 'option_name'.
func OptionByOptionName(db XODB, optionName string) (*Option, error) {
	v, found := c.Get("/options/option_name/" + optionName)
	if found {
		return v.(*Option), nil
	}
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`option_id, option_name, option_value, autoload ` +
		`FROM ` + GetTableName("options") + ` ` +
		`WHERE option_name = ?`

	// run query
	XOLog(sqlstr, optionName)
	o := Option{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, optionName).Scan(&o.OptionID, &o.OptionName, &o.OptionValue, &o.Autoload)
	if err != nil {
		return nil, err
	}

	c.Set("/options/option_name/"+optionName, &o, cache.DefaultExpiration)

	return &o, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// OptionByOptionID retrieves a row from '` + GetTableName("options") + `' as a Option.
//
// Generated from index 'options_option_id_pkey'.
func OptionByOptionID(db XODB, optionID uint64) (*Option, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`option_id, option_name, option_value, autoload ` +
		`FROM ` + GetTableName("options") + ` ` +
		`WHERE option_id = ?`

	// run query
	XOLog(sqlstr, optionID)
	o := Option{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, optionID).Scan(&o.OptionID, &o.OptionName, &o.OptionValue, &o.Autoload)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
