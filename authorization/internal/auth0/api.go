package auth0

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paulocesarvasco/web-chat/authorization/internal/requests"
)

type API interface {
	RequestAccessToken(ctx context.Context) (string, error)
	ValidateToken(ctx context.Context, bearer string) (bool, error)
}

type api struct{}

func (a *api) RequestAccessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", os.Getenv("AUTH0_DOMAIN"))
	requestPayload := AccessTokenRequestPayload{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Audience:     os.Getenv("AUTH0_AUDIENCE"),
		GrantType:    "client_credentials",
	}
	req, err := requests.NewPostRequest(ctx, url, requestPayload)
	if err != nil {
		return "", err
	}
	req.AddContentTypeJSON()
	code, res := req.Execute()
	if code != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", res)
	}
	var tokenResponse AccessTokenResponsePayload
	err = json.Unmarshal(res, &tokenResponse)
	if err != nil {
		return "", err
	}
	return tokenResponse.AccessToken, nil
}

func (a *api) ValidateToken(ctx context.Context, bearer string) (bool, error) {
	url := fmt.Sprintf("https://%s/userinfo", os.Getenv("AUTH0_DOMAIN"))
	req, err := requests.NewGetRequest(ctx, url)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %v", err)
	}
	req.AddAuthorizationToken(bearer)
	code, res := req.Execute()
	if code != http.StatusOK {
		log.Printf("%s", res)
		return false, fmt.Errorf("%s", res)
	}
	log.Printf("%s", res)
	return true, nil
}

func New() API {
	return &api{}
}
