package main

import (
	"log"
	"net/http"
	"web-chat/chat/internal/chat"
	"web-chat/chat/internal/ws"
)

func main() {
	room := chat.NewChatRoom()
	go room.Run()

	fs := http.FileServer(http.Dir("./clients"))
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
