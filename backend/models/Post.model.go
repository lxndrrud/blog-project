package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id          int64          `db:"post_id" json:"id"`
	IdAuthor    int64          `db:"post_id_author" json:"id_author"`
	AuthorLogin string         `db:"post_author_login" json:"author_login"`
	Title       string         `db:"post_title" json:"title"`
	Annotation  string         `db:"post_annotation" json:"annotation"`
	Picture     sql.NullString `db:"post_picture" json:"picture"`
	Text        string         `db:"post_text" json:"text"`
	Views       uint64         `db:"post_views" json:"views"`
	Approved    bool           `db:"post_approved" json:"approved"`
	Rejected    bool           `db:"post_rejected" json:"rejected"`
	TimeStart   sql.NullTime   `db:"post_time_start" json:"time_start"`
	TimeEnd     sql.NullTime   `db:"post_time_end" json:"time_end"`
	CreatedAt   time.Time      `db:"post_created_at" json:"created_at"`
}
