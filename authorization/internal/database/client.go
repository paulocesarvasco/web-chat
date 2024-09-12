package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateMySQLConnection() (*sql.DB, error) {
	user := "user"
	pass := "pass"
	mysqlHost := "mysql"
	mysqlPort := "3306"
	mysqlDatabase := "test_db"

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, mysqlHost, mysqlPort, mysqlDatabase)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to the MySQL database")
	return db, nil
}
