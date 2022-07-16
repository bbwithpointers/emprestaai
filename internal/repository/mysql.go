package repository

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// TODO: persistir
func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/meemprestaai")
	if err != nil {
		panic(err)
	}
	// important for mysql
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
