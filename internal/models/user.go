package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"partner.portal/internal/database"
	"partner.portal/internal/dto"
	"time"
)

func GetUserByID(id int64) (dto.User, error) {
	User := dto.User{}
	db := database.GetInstance()
	cand := db.QueryRow(`SELECT * FROM users WHERE id = $1`, id)
	err := cand.Scan(&User)
	if err != nil {
		return User, err
	}
	return User, nil
}

func GetUserByLogin(login string) (dto.User, error) {
	User := dto.User{}
	db := database.GetInstance()
	cand := db.QueryRow(`SELECT * FROM users WHERE login = $1`, login)
	err := cand.Scan(&User)
	if err != nil {
		return User, err
	}
	return User, nil
}

func UpdateUser(u dto.User) (sql.Result, error) {
	db := database.GetInstance()
	r, err := db.Exec(`UPDATE users SET login = $1, last_login = $2 WHERE id = $3`, u.Login, u.LastLogin, u.ID)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func CreateUser(u dto.User) (sql.Result, error) {
	var err error
	var h []byte
	var r sql.Result
	h, err = bcrypt.GenerateFromPassword([]byte(u.Pass), 12)
	if err != nil {
		return nil, err
	}
	db := database.GetInstance()
	r, err = db.Exec(`INSERT INTO users (login, pass, created_at) VALUES ($1, $2, $3)`, u.Login, string(h), time.Now())
	if err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteUser(id int64) (sql.Result, error) {
	db := database.GetInstance()
	r, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return r, err
	}
	return r, nil
}

func UpdateUserPass(u dto.User) (sql.Result, error) {
	var err error
	var h []byte
	var r sql.Result
	h, err = bcrypt.GenerateFromPassword([]byte(u.Pass), 12)
	if err != nil {
		return nil, err
	}
	db := database.GetInstance()
	r, err = db.Exec(`UPDATE users SET pass = $1 WHERE id = $2`, string(h), u.ID)
	if err != nil {
		return r, err
	}
	return r, nil
}
