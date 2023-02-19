package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	// dbHost := "192.168.14.2"
	dbUser := "root"
	dbPass := "MMp9ug6e"
	dbName := "manup-master"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}
