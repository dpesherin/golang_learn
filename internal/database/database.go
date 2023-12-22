package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"partner.portal/internal/config"
)

var instanceDb *sql.DB

func GetInstance() *sql.DB {
	var err error
	if instanceDb == nil {
		cfg := config.MustLoad()
		conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
		fmt.Println(conStr)
		instanceDb, err = sql.Open("postgres", conStr)
		if err != nil {
			log.Fatal(err)
		}
		err = instanceDb.Ping()
		if err != nil {
			panic(err)
		}
		fmt.Println("Connection to PostgreSQL: success")
	}
	return instanceDb
}
