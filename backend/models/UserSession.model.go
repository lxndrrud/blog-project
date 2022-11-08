package models

type UserSession struct {
	IdUser int64  `json:"session_id_user"`
	Token  string `json:"session_token"`
}
