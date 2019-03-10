package self

// xo "mysql://del:mysql_del@116.62.230.114:3306/del?parseTime=true" -o wpa --template-path ./templates

// PostmetaByPostIDMetaKey retrieves a row from 'del.postmeta' as a Postmetum.
//
// Generated from index 'postmeta_post_id_meta_key_pkey'.
func PostmetaByPostIDMetaKey(db XODB, postID uint64, metaKey string) (*Postmetum, error) {
	var err error

	// sql query
	var sqlstr = `SELECT ` +
		`meta_id, post_id, meta_key, meta_value ` +
		`FROM ` + GetTableName("postmeta") + ` ` +
		`WHERE post_id = ? AND meta_key = ?`

	// run query
	XOLog(sqlstr, postID, metaKey)
	p := Postmetum{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, postID, metaKey).Scan(&p.MetaID, &p.PostID, &p.MetaKey, &p.MetaValue)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
