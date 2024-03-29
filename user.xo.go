// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// User represents a row from '` + GetTableName("users") + `'.
type User struct {
	ID                uint64    `json:"ID"`                  // ID
	UserLogin         string    `json:"user_login"`          // user_login
	UserPass          string    `json:"user_pass"`           // user_pass
	UserNicename      string    `json:"user_nicename"`       // user_nicename
	UserEmail         string    `json:"user_email"`          // user_email
	UserURL           string    `json:"user_url"`            // user_url
	UserRegistered    time.Time `json:"user_registered"`     // user_registered
	UserActivationKey string    `json:"user_activation_key"` // user_activation_key
	UserStatus        int       `json:"user_status"`         // user_status
	DisplayName       string    `json:"display_name"`        // display_name
	Spam              int8      `json:"spam"`                // spam
	Deleted           int8      `json:"deleted"`             // deleted

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deled provides information if the User has been deled from the database.
func (u *User) Deled() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("users") + ` (` +
		`user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name, spam, deleted` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.UserLogin, u.UserPass, u.UserNicename, u.UserEmail, u.UserURL, u.UserRegistered, u.UserActivationKey, u.UserStatus, u.DisplayName, u.Spam, u.Deleted)
	res, err := db.Exec(sqlstr, u.UserLogin, u.UserPass, u.UserNicename, u.UserEmail, u.UserURL, u.UserRegistered, u.UserActivationKey, u.UserStatus, u.DisplayName, u.Spam, u.Deleted)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = uint64(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("users") + ` SET ` +
		`user_login = ?, user_pass = ?, user_nicename = ?, user_email = ?, user_url = ?, user_registered = ?, user_activation_key = ?, user_status = ?, display_name = ?, spam = ?, deleted = ?` +
		` WHERE ID = ?`

	// run query
	XOLog(sqlstr, u.UserLogin, u.UserPass, u.UserNicename, u.UserEmail, u.UserURL, u.UserRegistered, u.UserActivationKey, u.UserStatus, u.DisplayName, u.Spam, u.Deleted, u.ID)
	_, err = db.Exec(sqlstr, u.UserLogin, u.UserPass, u.UserNicename, u.UserEmail, u.UserURL, u.UserRegistered, u.UserActivationKey, u.UserStatus, u.DisplayName, u.Spam, u.Deleted, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("users") + ` WHERE ID = ?`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// UsersByUserEmail retrieves a row from '` + GetTableName("users") + `' as a User.
//
// Generated from index 'user_email'.
func UsersByUserEmail(db XODB, userEmail string) ([]*User, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name, spam, deleted ` +
		`FROM ` + GetTableName("users") + ` ` +
		`WHERE user_email = ?`

	// run query
	XOLog(sqlstr, userEmail)
	q, err := db.Query(sqlstr, userEmail)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*User{}
	for q.Next() {
		u := User{
			_exists: true,
		}

		// scan
		err = q.Scan(&u.ID, &u.UserLogin, &u.UserPass, &u.UserNicename, &u.UserEmail, &u.UserURL, &u.UserRegistered, &u.UserActivationKey, &u.UserStatus, &u.DisplayName, &u.Spam, &u.Deleted)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// UsersByUserLogin retrieves a row from '` + GetTableName("users") + `' as a User.
//
// Generated from index 'user_login_key'.
func UsersByUserLogin(db XODB, userLogin string) ([]*User, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name, spam, deleted ` +
		`FROM ` + GetTableName("users") + ` ` +
		`WHERE user_login = ?`

	// run query
	XOLog(sqlstr, userLogin)
	q, err := db.Query(sqlstr, userLogin)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*User{}
	for q.Next() {
		u := User{
			_exists: true,
		}

		// scan
		err = q.Scan(&u.ID, &u.UserLogin, &u.UserPass, &u.UserNicename, &u.UserEmail, &u.UserURL, &u.UserRegistered, &u.UserActivationKey, &u.UserStatus, &u.DisplayName, &u.Spam, &u.Deleted)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// UsersByUserNicename retrieves a row from '` + GetTableName("users") + `' as a User.
//
// Generated from index 'user_nicename'.
func UsersByUserNicename(db XODB, userNicename string) ([]*User, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name, spam, deleted ` +
		`FROM ` + GetTableName("users") + ` ` +
		`WHERE user_nicename = ?`

	// run query
	XOLog(sqlstr, userNicename)
	q, err := db.Query(sqlstr, userNicename)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*User{}
	for q.Next() {
		u := User{
			_exists: true,
		}

		// scan
		err = q.Scan(&u.ID, &u.UserLogin, &u.UserPass, &u.UserNicename, &u.UserEmail, &u.UserURL, &u.UserRegistered, &u.UserActivationKey, &u.UserStatus, &u.DisplayName, &u.Spam, &u.Deleted)
		if err != nil {
			return nil, err
		}

		res = append(res, &u)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// UserByID retrieves a row from '` + GetTableName("users") + `' as a User.
//
// Generated from index 'users_ID_pkey'.
func UserByID(db XODB, id uint64) (*User, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name, spam, deleted ` +
		`FROM ` + GetTableName("users") + ` ` +
		`WHERE ID = ?`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.UserLogin, &u.UserPass, &u.UserNicename, &u.UserEmail, &u.UserURL, &u.UserRegistered, &u.UserActivationKey, &u.UserStatus, &u.DisplayName, &u.Spam, &u.Deleted)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
