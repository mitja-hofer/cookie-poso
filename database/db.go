package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetConnection() *sql.DB {
	// TODO: ssl?
	cfg := mysql.Config{
		Addr:                 os.Getenv("COOKIE_DB_HOST"),
		User:                 os.Getenv("COOKIE_DB_USER"),
		Passwd:               os.Getenv("COOKIE_DB_PASSWORD"),
		Net:                  "tcp",
		DBName:               "cookie",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
