package models

type Permission struct {
	Id    int64  `db:"permission_id" json:"-"`
	Title string `db:"permission_title" json:"-"`
	Code  string `db:"permission_code" json:"title"`
}
