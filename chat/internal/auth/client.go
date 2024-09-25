package auth

import (
	"log"
	"net/http"
)

func ValidateCredentials(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if token := r.URL.Query().Get("token"); token != "" {
			next(w, r)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		authReq, err := http.NewRequest(http.MethodPost, "http://auth_service:8081/login", nil)
		if err != nil {
			http.Error(w, "Failed to create request to auth_service", http.StatusInternalServerError)
			return
		}
		authReq.Header.Set("Authorization", authHeader)

		client := &http.Client{}
		res, err := client.Do(authReq)
		if err != nil {
			log.Print(err)
			http.Error(w, "Authentication failed", http.StatusInternalServerError)
			return
		}
		if res.StatusCode != http.StatusOK {
			http.Error(w, "Authentication failed", http.StatusUnauthorized)
			return
		}
		token := res.Header.Get("token")
		w.Header().Set("Authorization", "Bearer "+token)
	}
}
