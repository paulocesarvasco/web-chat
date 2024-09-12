package main

import (
	"github.com/paulocesarvasco/web-chat/chat/internal/chat"
	"github.com/paulocesarvasco/web-chat/chat/internal/ws"
	"log"
	"net/http"
)

func main() {
	room := chat.NewChatRoom()
	go room.Run()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleConnections(w, r, room)
	})

	log.Println("Chat server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
