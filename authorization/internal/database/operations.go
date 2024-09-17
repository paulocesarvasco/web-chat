package database

import (
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
