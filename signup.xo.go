// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// Signup represents a row from '` + GetTableName("signups") + `'.
type Signup struct {
	SignupID      int64          `json:"signup_id"`      // signup_id
	Domain        string         `json:"domain"`         // domain
	Path          string         `json:"path"`           // path
	Title         string         `json:"title"`          // title
	UserLogin     string         `json:"user_login"`     // user_login
	UserEmail     string         `json:"user_email"`     // user_email
	Registered    time.Time      `json:"registered"`     // registered
	Activated     time.Time      `json:"activated"`      // activated
	Active        bool           `json:"active"`         // active
	ActivationKey string         `json:"activation_key"` // activation_key
	Meta          sql.NullString `json:"meta"`           // meta

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Signup exists in the database.
func (s *Signup) Exists() bool {
	return s._exists
}

// Deled provides information if the Signup has been deled from the database.
func (s *Signup) Deled() bool {
	return s._deleted
}

// Insert inserts the Signup to the database.
func (s *Signup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if s._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("signups") + ` (` +
		`domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, s.Domain, s.Path, s.Title, s.UserLogin, s.UserEmail, s.Registered, s.Activated, s.Active, s.ActivationKey, s.Meta)
	res, err := db.Exec(sqlstr, s.Domain, s.Path, s.Title, s.UserLogin, s.UserEmail, s.Registered, s.Activated, s.Active, s.ActivationKey, s.Meta)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	s.SignupID = int64(id)
	s._exists = true

	return nil
}

// Update updates the Signup in the database.
func (s *Signup) Update(db XODB) error {
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
	var sqlstr = `UPDATE ` + GetTableName("signups") + ` SET ` +
		`domain = ?, path = ?, title = ?, user_login = ?, user_email = ?, registered = ?, activated = ?, active = ?, activation_key = ?, meta = ?` +
		` WHERE signup_id = ?`

	// run query
	XOLog(sqlstr, s.Domain, s.Path, s.Title, s.UserLogin, s.UserEmail, s.Registered, s.Activated, s.Active, s.ActivationKey, s.Meta, s.SignupID)
	_, err = db.Exec(sqlstr, s.Domain, s.Path, s.Title, s.UserLogin, s.UserEmail, s.Registered, s.Activated, s.Active, s.ActivationKey, s.Meta, s.SignupID)
	return err
}

// Save saves the Signup to the database.
func (s *Signup) Save(db XODB) error {
	if s.Exists() {
		return s.Update(db)
	}

	return s.Insert(db)
}

// Delete deletes the Signup from the database.
func (s *Signup) Delete(db XODB) error {
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
	var sqlstr = `DELETE FROM ` + GetTableName("signups") + ` WHERE signup_id = ?`

	// run query
	XOLog(sqlstr, s.SignupID)
	_, err = db.Exec(sqlstr, s.SignupID)
	if err != nil {
		return err
	}

	// set deleted
	s._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SignupsByActivationKey retrieves a row from '` + GetTableName("signups") + `' as a Signup.
//
// Generated from index 'activation_key'.
func SignupsByActivationKey(db XODB, activationKey string) ([]*Signup, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`signup_id, domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta ` +
		`FROM ` + GetTableName("signups") + ` ` +
		`WHERE activation_key = ?`

	// run query
	XOLog(sqlstr, activationKey)
	q, err := db.Query(sqlstr, activationKey)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Signup{}
	for q.Next() {
		s := Signup{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.SignupID, &s.Domain, &s.Path, &s.Title, &s.UserLogin, &s.UserEmail, &s.Registered, &s.Activated, &s.Active, &s.ActivationKey, &s.Meta)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SignupsByDomainPath retrieves a row from '` + GetTableName("signups") + `' as a Signup.
//
// Generated from index 'domain_path'.
func SignupsByDomainPath(db XODB, domain string, path string) ([]*Signup, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`signup_id, domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta ` +
		`FROM ` + GetTableName("signups") + ` ` +
		`WHERE domain = ? AND path = ?`

	// run query
	XOLog(sqlstr, domain, path)
	q, err := db.Query(sqlstr, domain, path)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Signup{}
	for q.Next() {
		s := Signup{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.SignupID, &s.Domain, &s.Path, &s.Title, &s.UserLogin, &s.UserEmail, &s.Registered, &s.Activated, &s.Active, &s.ActivationKey, &s.Meta)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SignupBySignupID retrieves a row from '` + GetTableName("signups") + `' as a Signup.
//
// Generated from index 'signups_signup_id_pkey'.
func SignupBySignupID(db XODB, signupID int64) (*Signup, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`signup_id, domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta ` +
		`FROM ` + GetTableName("signups") + ` ` +
		`WHERE signup_id = ?`

	// run query
	XOLog(sqlstr, signupID)
	s := Signup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, signupID).Scan(&s.SignupID, &s.Domain, &s.Path, &s.Title, &s.UserLogin, &s.UserEmail, &s.Registered, &s.Activated, &s.Active, &s.ActivationKey, &s.Meta)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SignupsByUserEmail retrieves a row from '` + GetTableName("signups") + `' as a Signup.
//
// Generated from index 'user_email'.
func SignupsByUserEmail(db XODB, userEmail string) ([]*Signup, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`signup_id, domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta ` +
		`FROM ` + GetTableName("signups") + ` ` +
		`WHERE user_email = ?`

	// run query
	XOLog(sqlstr, userEmail)
	q, err := db.Query(sqlstr, userEmail)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Signup{}
	for q.Next() {
		s := Signup{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.SignupID, &s.Domain, &s.Path, &s.Title, &s.UserLogin, &s.UserEmail, &s.Registered, &s.Activated, &s.Active, &s.ActivationKey, &s.Meta)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// SignupsByUserLoginUserEmail retrieves a row from '` + GetTableName("signups") + `' as a Signup.
//
// Generated from index 'user_login_email'.
func SignupsByUserLoginUserEmail(db XODB, userLogin string, userEmail string) ([]*Signup, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`signup_id, domain, path, title, user_login, user_email, registered, activated, active, activation_key, meta ` +
		`FROM ` + GetTableName("signups") + ` ` +
		`WHERE user_login = ? AND user_email = ?`

	// run query
	XOLog(sqlstr, userLogin, userEmail)
	q, err := db.Query(sqlstr, userLogin, userEmail)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Signup{}
	for q.Next() {
		s := Signup{
			_exists: true,
		}

		// scan
		err = q.Scan(&s.SignupID, &s.Domain, &s.Path, &s.Title, &s.UserLogin, &s.UserEmail, &s.Registered, &s.Activated, &s.Active, &s.ActivationKey, &s.Meta)
		if err != nil {
			return nil, err
		}

		res = append(res, &s)
	}

	return res, nil
}
