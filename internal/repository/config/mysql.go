package config

import (
	"database/sql"
	"github.com/brunogbarros/emprestaai.git/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

// ConnectDB : Creates the connection to db with string connection
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnDB)
	if err != nil {
		log.Fatal("Error while trying to connect to DB")
	}
	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
