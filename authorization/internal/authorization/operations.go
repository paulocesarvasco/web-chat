package authorization

import (
	"fmt"

	"github.com/paulocesarvasco/web-chat/authorization/internal/database"
	"github.com/paulocesarvasco/web-chat/authorization/internal/resources"
)

type auth struct {
	db database.Client
}

func (a *auth) CreateNewClient(clientInfo resources.Client) error {
	hashedPassword, err := hashPassword(clientInfo.Password)
	if err != nil {
		return fmt.Errorf("failed to  hash password: %v", err)
	}
	clientInfo.Password = hashedPassword
	return a.db.CreateClient(clientInfo)
}

func (a *auth) CheckClientPassword(username, password string) (bool, error) {
	client, err := a.db.GetClientByUsername(username)
	if err != nil {
		return false, err
	}
	return verifyPassword(client.Password, password)
}
