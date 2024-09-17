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

	hashedPassword, err := hashPassword(client.Password)
	if err != nil {
		return fmt.Errorf("failed to  hash password: %v", err)
	}

	_, err = stmt.Exec(client.Email, hashedPassword, client.Name, client.Username)
	if err != nil {
		return fmt.Errorf("failed to insert client: %v", err)
	}
	return nil
}
