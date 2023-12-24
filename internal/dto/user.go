package dto

import "time"

type User struct {
	ID        int64     `json:"id"`
	Login     string    `json:"login"`
	Pass      string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
	Perm      `json:"perm"`
}

type Perm struct {
	Role string `json:"role"`
}
