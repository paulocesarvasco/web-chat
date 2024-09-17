package database

import (
	"database/sql"
	"fmt"

	"github.com/paulocesarvasco/web-chat/authorization/internal/resources"
)

func (c *client) CreateClient(client resources.Client) error {
	stmt, err := c.db.Prepare("INSERT INTO users (email, password, name, username) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.Email, client.Password, client.Name, client.Username)
	if err != nil {
		return fmt.Errorf("failed to insert client: %v", err)
	}
	return nil
}

func (c *client) GetClientByUsername(username string) (resources.Client, error) {
	var client resources.Client
	query := "SELECT email, name, password, username FROM users WHERE username = ?"
	err := c.db.QueryRow(query, username).Scan(&client.Email, &client.Name, &client.Password, &client.Username)
	switch err {
	case nil:
		return client, nil
	case sql.ErrNoRows:
		return resources.Client{}, fmt.Errorf("usuário não encontrado")
	default:
		return resources.Client{}, fmt.Errorf("erro ao buscar o usuário: %v", err)
	}
}
