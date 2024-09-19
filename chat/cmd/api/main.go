package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/paulocesarvasco/web-chat/chat/internal/auth"
	"github.com/paulocesarvasco/web-chat/chat/internal/chat"
	"github.com/paulocesarvasco/web-chat/chat/internal/ws"
)

func main() {
	room := chat.NewChatRoom()
	go room.Run()

	r := chi.NewMux()

	fs := http.FileServer(http.Dir("/static"))
	r.Handle("/*", fs)
	r.HandleFunc("/chat", auth.ValidateCredentials(
		func(w http.ResponseWriter, r *http.Request) {
			ws.HandleConnections(w, r, room)
		}))

	log.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
