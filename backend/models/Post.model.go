package models

import (
	"database/sql"
)

type Post struct {
	Id          int64        `db:"post_id" json:"id"`
	IdAuthor    int64        `db:"post_id_author" json:"id_author"`
	AuthorLogin string       `db:"post_author_login" json:"author_login"`
	Title       string       `db:"post_title" json:"title"`
	Text        string       `db:"post_text" json:"text"`
	Views       uint64       `db:"post_views" json:"views"`
	Approved    bool         `db:"post_approved" json:"approved"`
	TimeEnd     sql.NullTime `db:"post_time_end" json:"time_end"`
}
