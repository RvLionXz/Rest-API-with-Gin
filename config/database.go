package config

import (
	"database/sql"
	"fmt"

	// Gunakan underscore _ karena kita hanya butuh "efek samping" registrasi driver
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (*sql.DB, error) {
	username := "root"
	password := "reval123"
	hostname := "127.0.0.1"
	port := "3306"
	dbname := "task_manager_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostname, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
