package models

import "encoding/json"

type User struct {
	Id       int64  `db:"user_id"`
	Login    string `db:"user_login"`
	Password string `db:"user_password"`
}

func (c User) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":    c.Id,
		"login": c.Login,
	}
	return json.Marshal(m)
}
