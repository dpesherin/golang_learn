package models

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"partner.portal/internal/database"
	"partner.portal/internal/dto"
	"time"
)

type MyCustomClaims struct {
	UserID int64
	jwt.StandardClaims
}

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
	err := cand.Scan(&User.ID, &User.Login, &User.Pass, &User.CreatedAt, &User.LastLogin)
	if err != nil {
		return User, err
	}
	return User, nil
}

func UpdateUser(u dto.User) (string, error) {
	db := database.GetInstance()
	_, err := db.Exec(`UPDATE users SET login = $1, last_login = $2 WHERE id = $3`, u.Login, u.LastLogin, u.ID)
	if err != nil {
		return "Database error", err
	}
	return "Update: Success", nil
}

func CreateUser(u dto.User) (string, error) {
	var err error
	var h []byte
	h, err = bcrypt.GenerateFromPassword([]byte(u.Pass), 12)
	if err != nil {
		return "Crypt: Error", err
	}
	db := database.GetInstance()
	_, err = db.Exec(`INSERT INTO users (login, pass, created_at) VALUES ($1, $2, $3)`, u.Login, string(h), time.Now())
	if err != nil {
		return "Database error", err
	}
	return "Create: Success", nil
}

func DeleteUser(id int64) (string, error) {
	db := database.GetInstance()
	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return "Database error", err
	}
	return "Delete: Success", nil
}

func UpdateUserPass(u dto.User) (string, error) {
	var err error
	var h []byte
	h, err = bcrypt.GenerateFromPassword([]byte(u.Pass), 12)
	if err != nil {
		return "Crypt: Error", err
	}
	db := database.GetInstance()
	_, err = db.Exec(`UPDATE users SET pass = $1 WHERE id = $2`, string(h), u.ID)
	if err != nil {
		return "Database Error", err
	}
	return "UpdatePass: Success", nil
}

func GenJWT(u dto.User, expHours time.Duration) (string, error) {
	secret := []byte("SeCr3tK3y")
	expirationTime := time.Now().Add(expHours * time.Hour)

	claims := MyCustomClaims{
		u.ID,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "Application",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "JWT Error", err
	}
	return tokenString, nil
}
