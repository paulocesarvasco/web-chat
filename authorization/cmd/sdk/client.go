package sdk

import (
	"context"
	"log"
	"net/http"

	"github.com/paulocesarvasco/web-chat/authorization/internal/requests"
)

const AUTH_HOST = "http://127.0.0.1:8081/login"

type client struct{}

type Client interface {
	Login(ctx context.Context, user, pass string) bool
}

func NewAuthClient() Client {
	return &client{}
}

func (c *client) Login(ctx context.Context, user, pass string) bool {
	req, err := requests.NewPostRequest(ctx, AUTH_HOST, []byte("{}"))
	if err != nil {
		log.Print(err)
		return false
	}
	code, _ := req.Execute()
	return code == http.StatusOK
}
