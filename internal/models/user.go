package models

type User struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Perm
}

type Perm struct {
	Role string `json:"role"`
}
