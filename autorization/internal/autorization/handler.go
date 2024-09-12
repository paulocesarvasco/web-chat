package autorization

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func ValidateCredentials(w http.ResponseWriter, r *http.Request) {
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
	log.Println(username)
	log.Println(password)
}
