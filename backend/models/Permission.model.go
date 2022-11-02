package models

type Permission struct {
	Id    int64  `db:"permission_id"`
	Title string `db:"permission_title"`
	Code  string `db:"permission_code"`
}
