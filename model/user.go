package model

type UserData struct {
	UserID int64  `db:"user_id"json:"user_id"`
	Name   string `db:"name"json:"name"`
}
