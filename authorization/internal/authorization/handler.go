package authorization

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func (a *api) ValidateCredentials(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Failed to validate credentials", http.StatusInternalServerError)
		return
	}
	if !isValid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
