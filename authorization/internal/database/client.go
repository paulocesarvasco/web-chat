package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Client interface {
}

type client struct {
	db *sql.DB
}

func NewClient() (Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Successfully connected to the MySQL database")
	return &client{db: db}, nil
}
