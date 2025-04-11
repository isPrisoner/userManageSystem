package utils

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:lgdzz666@tcp(127.0.0.1:3306)/user_admin")
	if err != nil {
		log.Println(err)
		return nil
	}
	return db
}
