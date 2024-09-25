package authorization

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/paulocesarvasco/web-chat/authorization/internal/resources"
)

func (a *api) ValidateCredentials() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(auth, "Basic ") {
			http.Error(w, "Invalid Authorization", http.StatusBadRequest)
			return
		}
		payload, err := base64.StdEncoding.DecodeString(auth[len("Basic "):])
		if err != nil {
			http.Error(w, "Invalid base64 encoding", http.StatusBadRequest)
			return
		}
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			http.Error(w, "Invalid Authorization format", http.StatusBadRequest)
			return
		}
		username, password := pair[0], pair[1]
		isValid, err := a.CheckClientPassword(username, password)
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to validate credentials", http.StatusInternalServerError)
			return
		}
		if !isValid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
		token, err := a.RequestAccessToken(r.Context())
		if err != nil {
			log.Print("failed request access token", err)
			http.Error(w, "failed request access token", http.StatusInternalServerError)
		}
		w.Header().Add("token", token)
	}
}

func (a *api) ValidateToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missed token", http.StatusUnauthorized)
			return
		}
		isValid, err := a.API.ValidateToken(r.Context(), token)
		if err != nil {
			http.Error(w, "failed to validate token", http.StatusInternalServerError)
			return
		}
		if !isValid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
	}
}

func (a *api) CreateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientInfo resources.Client
		err := json.NewDecoder(r.Body).Decode(&clientInfo)
		if err != nil {
			http.Error(w, "failed to decode payload", http.StatusInternalServerError)
			return
		}
		err = a.CreateNewClient(clientInfo)
		if err != nil {
			http.Error(w, "failed to create client", http.StatusInternalServerError)
			return
		}
	}
}
