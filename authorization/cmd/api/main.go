package main

import (
	"github.com/paulocesarvasco/web-chat/authorization/internal/authorization"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		authorization.ValidateCredentials(w, r)
	})
	log.Println("Auth service started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
