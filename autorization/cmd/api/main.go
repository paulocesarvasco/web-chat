package main

import (
	"github.com/paulocesarvasco/web-chat/autorization/internal/autorization"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		autorization.ValidateCredentials(w, r)
	})
	log.Println("Auth service started on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
