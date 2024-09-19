package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/paulocesarvasco/web-chat/authorization/internal/authorization"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Auth service started on :8081")
	err = http.ListenAndServe(":8081", authorization.NewAPI().Engine())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
