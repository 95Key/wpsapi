// Package wpsapi contains the types for schema 'del'.
// i say package is wpsapi
package wpsapi

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"

	"util"
)

// Post represents a row from '` + GetTableName("posts") + `'.
type Post struct {
	ID                  uint64    `json:"ID"`                    // ID
	PostAuthor          uint64    `json:"post_author"`           // post_author
	PostDate            time.Time `json:"post_date"`             // post_date
	PostDateGmt         time.Time `json:"post_date_gmt"`         // post_date_gmt
	PostContent         string    `json:"post_content"`          // post_content
	PostTitle           string    `json:"post_title"`            // post_title
	PostExcerpt         string    `json:"post_excerpt"`          // post_excerpt
	PostStatus          string    `json:"post_status"`           // post_status
	CommentStatus       string    `json:"comment_status"`        // comment_status
	PingStatus          string    `json:"ping_status"`           // ping_status
	PostPassword        string    `json:"post_password"`         // post_password
	PostName            string    `json:"post_name"`             // post_name
	ToPing              string    `json:"to_ping"`               // to_ping
	Pinged              string    `json:"pinged"`                // pinged
	PostModified        time.Time `json:"post_modified"`         // post_modified
	PostModifiedGmt     time.Time `json:"post_modified_gmt"`     // post_modified_gmt
	PostContentFiltered string    `json:"post_content_filtered"` // post_content_filtered
	PostParent          uint64    `json:"post_parent"`           // post_parent
	GUID                string    `json:"guid"`                  // guid
	MenuOrder           int       `json:"menu_order"`            // menu_order
	PostType            string    `json:"post_type"`             // post_type
	PostMimeType        string    `json:"post_mime_type"`        // post_mime_type
	CommentCount        int64     `json:"comment_count"`         // comment_count

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Post exists in the database.
func (p *Post) Exists() bool {
	return p._exists
}

// Deled provides information if the Post has been deled from the database.
func (p *Post) Deled() bool {
	return p._deleted
}

// Insert inserts the Post to the database.
func (p *Post) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	var sqlstr = `INSERT INTO ` + GetTableName("posts") + ` (` +
		`post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, p.PostAuthor, p.PostDate, p.PostDateGmt, p.PostContent, p.PostTitle, p.PostExcerpt, p.PostStatus, p.CommentStatus, p.PingStatus, p.PostPassword, p.PostName, p.ToPing, p.Pinged, p.PostModified, p.PostModifiedGmt, p.PostContentFiltered, p.PostParent, p.GUID, p.MenuOrder, p.PostType, p.PostMimeType, p.CommentCount)
	res, err := db.Exec(sqlstr, p.PostAuthor, p.PostDate, p.PostDateGmt, p.PostContent, p.PostTitle, p.PostExcerpt, p.PostStatus, p.CommentStatus, p.PingStatus, p.PostPassword, p.PostName, p.ToPing, p.Pinged, p.PostModified, p.PostModifiedGmt, p.PostContentFiltered, p.PostParent, p.GUID, p.MenuOrder, p.PostType, p.PostMimeType, p.CommentCount)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	p.ID = uint64(id)
	p._exists = true

	return nil
}

// Update updates the Post in the database.
func (p *Post) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	var sqlstr = `UPDATE ` + GetTableName("posts") + ` SET ` +
		`post_author = ?, post_date = ?, post_date_gmt = ?, post_content = ?, post_title = ?, post_excerpt = ?, post_status = ?, comment_status = ?, ping_status = ?, post_password = ?, post_name = ?, to_ping = ?, pinged = ?, post_modified = ?, post_modified_gmt = ?, post_content_filtered = ?, post_parent = ?, guid = ?, menu_order = ?, post_type = ?, post_mime_type = ?, comment_count = ?` +
		` WHERE ID = ?`

	// run query
	XOLog(sqlstr, p.PostAuthor, p.PostDate, p.PostDateGmt, p.PostContent, p.PostTitle, p.PostExcerpt, p.PostStatus, p.CommentStatus, p.PingStatus, p.PostPassword, p.PostName, p.ToPing, p.Pinged, p.PostModified, p.PostModifiedGmt, p.PostContentFiltered, p.PostParent, p.GUID, p.MenuOrder, p.PostType, p.PostMimeType, p.CommentCount, p.ID)
	_, err = db.Exec(sqlstr, p.PostAuthor, p.PostDate, p.PostDateGmt, p.PostContent, p.PostTitle, p.PostExcerpt, p.PostStatus, p.CommentStatus, p.PingStatus, p.PostPassword, p.PostName, p.ToPing, p.Pinged, p.PostModified, p.PostModifiedGmt, p.PostContentFiltered, p.PostParent, p.GUID, p.MenuOrder, p.PostType, p.PostMimeType, p.CommentCount, p.ID)
	return err
}

// Save saves the Post to the database.
func (p *Post) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Delete deletes the Post from the database.
func (p *Post) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	var sqlstr = `DELETE FROM ` + GetTableName("posts") + ` WHERE ID = ?`

	// run query
	XOLog(sqlstr, p.ID)
	_, err = db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// PostsByPostAuthor retrieves a row from '` + GetTableName("posts") + `' as a Post.
//
// Generated from index 'post_author'.
func PostsByPostAuthor(db XODB, postAuthor uint64) ([]*Post, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count ` +
		`FROM ` + GetTableName("posts") + ` ` +
		`WHERE post_author = ?`

	// run query
	XOLog(sqlstr, postAuthor)
	q, err := db.Query(sqlstr, postAuthor)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Post{}
	for q.Next() {
		p := Post{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.PostAuthor, &p.PostDate, &p.PostDateGmt, &p.PostContent, &p.PostTitle, &p.PostExcerpt, &p.PostStatus, &p.CommentStatus, &p.PingStatus, &p.PostPassword, &p.PostName, &p.ToPing, &p.Pinged, &p.PostModified, &p.PostModifiedGmt, &p.PostContentFiltered, &p.PostParent, &p.GUID, &p.MenuOrder, &p.PostType, &p.PostMimeType, &p.CommentCount)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// PostsByPostName retrieves a row from '` + GetTableName("posts") + `' as a Post.
//
// Generated from index 'post_name'.
func PostsByPostName(db XODB, postName string) ([]*Post, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count ` +
		`FROM ` + GetTableName("posts") + ` ` +
		`WHERE post_name = ?`

	// run query
	XOLog(sqlstr, postName)
	q, err := db.Query(sqlstr, postName)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Post{}
	for q.Next() {
		p := Post{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.PostAuthor, &p.PostDate, &p.PostDateGmt, &p.PostContent, &p.PostTitle, &p.PostExcerpt, &p.PostStatus, &p.CommentStatus, &p.PingStatus, &p.PostPassword, &p.PostName, &p.ToPing, &p.Pinged, &p.PostModified, &p.PostModifiedGmt, &p.PostContentFiltered, &p.PostParent, &p.GUID, &p.MenuOrder, &p.PostType, &p.PostMimeType, &p.CommentCount)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// PostsByPostParent retrieves a row from '` + GetTableName("posts") + `' as a Post.
//
// Generated from index 'post_parent'.
func PostsByPostParent(db XODB, postParent uint64) ([]*Post, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count ` +
		`FROM ` + GetTableName("posts") + ` ` +
		`WHERE post_parent = ?`

	// run query
	XOLog(sqlstr, postParent)
	q, err := db.Query(sqlstr, postParent)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Post{}
	for q.Next() {
		p := Post{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.PostAuthor, &p.PostDate, &p.PostDateGmt, &p.PostContent, &p.PostTitle, &p.PostExcerpt, &p.PostStatus, &p.CommentStatus, &p.PingStatus, &p.PostPassword, &p.PostName, &p.ToPing, &p.Pinged, &p.PostModified, &p.PostModifiedGmt, &p.PostContentFiltered, &p.PostParent, &p.GUID, &p.MenuOrder, &p.PostType, &p.PostMimeType, &p.CommentCount)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// PostByID retrieves a row from '` + GetTableName("posts") + `' as a Post.
//
// Generated from index 'posts_ID_pkey'.
func PostByID(db XODB, id uint64) (*Post, error) {

	v, found := c.Get("/post/ID/" + util.Uint64ToString(id))
	if found {

		return v.(*Post), nil
	}

	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count ` +
		`FROM ` + GetTableName("posts") + ` ` +
		`WHERE ID = ?`

	// run query
	XOLog(sqlstr, id)
	p := Post{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.PostAuthor, &p.PostDate, &p.PostDateGmt, &p.PostContent, &p.PostTitle, &p.PostExcerpt, &p.PostStatus, &p.CommentStatus, &p.PingStatus, &p.PostPassword, &p.PostName, &p.ToPing, &p.Pinged, &p.PostModified, &p.PostModifiedGmt, &p.PostContentFiltered, &p.PostParent, &p.GUID, &p.MenuOrder, &p.PostType, &p.PostMimeType, &p.CommentCount)
	if err != nil {
		return nil, err
	}
	c.Set("/post/ID/"+util.Uint64ToString(id), &p, cache.DefaultExpiration)
	return &p, nil
}

//{ {- $table := (schema .Schema .Type.Table.TableName) -} }// PostsByPostTypePostStatusPostDateID retrieves a row from '` + GetTableName("posts") + `' as a Post.
//
// Generated from index 'type_status_date'.
func PostsByPostTypePostStatusPostDateID(db XODB, postType string, postStatus string, postDate time.Time, id uint64) ([]*Post, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`ID, post_author, post_date, post_date_gmt, post_content, post_title, post_excerpt, post_status, comment_status, ping_status, post_password, post_name, to_ping, pinged, post_modified, post_modified_gmt, post_content_filtered, post_parent, guid, menu_order, post_type, post_mime_type, comment_count ` +
		`FROM ` + GetTableName("posts") + ` ` +
		`WHERE post_type = ? AND post_status = ? AND post_date = ? AND ID = ?`

	// run query
	XOLog(sqlstr, postType, postStatus, postDate, id)
	q, err := db.Query(sqlstr, postType, postStatus, postDate, id)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Post{}
	for q.Next() {
		p := Post{
			_exists: true,
		}

		// scan
		err = q.Scan(&p.ID, &p.PostAuthor, &p.PostDate, &p.PostDateGmt, &p.PostContent, &p.PostTitle, &p.PostExcerpt, &p.PostStatus, &p.CommentStatus, &p.PingStatus, &p.PostPassword, &p.PostName, &p.ToPing, &p.Pinged, &p.PostModified, &p.PostModifiedGmt, &p.PostContentFiltered, &p.PostParent, &p.GUID, &p.MenuOrder, &p.PostType, &p.PostMimeType, &p.CommentCount)
		if err != nil {
			return nil, err
		}

		res = append(res, &p)
	}

	return res, nil
}