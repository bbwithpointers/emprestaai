package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// important for mysql
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
