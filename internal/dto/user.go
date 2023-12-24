package dto

import (
	"database/sql"
)

type User struct {
	ID        int64        `json:"id" sql:"id"`
	Login     string       `json:"login" sql:"login"`
	Pass      string       `json:"pass" sql:"pass"`
	CreatedAt sql.NullTime `json:"created_at" sql:"created_at"`
	LastLogin sql.NullTime `json:"last_login" sql:"last_login"`
	Perm      `json:"perm" sql:"perm"`
}

type Perm struct {
	Role string `json:"role" sql:"role"`
}
