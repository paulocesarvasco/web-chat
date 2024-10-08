package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/paulocesarvasco/web-chat/authorization/internal/resources"
)

type Client interface {
	CreateClient(resources.Client) error
	GetClientByUsername(username string) (resources.Client, error)
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
	return &client{db: db}, nil
}
