package ws

import (
	// "io"
	"log"
	"net/http"

	"github.com/paulocesarvasco/web-chat/chat/internal/chat"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request, room *chat.ChatRoom) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &chat.Client{
		Conn: ws,
		Send: make(chan []byte),
	}

	room.Register <- client

	go client.ReadMessages(room)
	go client.WriteMessages()
}
