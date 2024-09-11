package ws

import (
	"log"
	"net/http"
	"web-chat/internal/chat"

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

	// Ensure that the client is unregistered when they disconnect
	defer func() {
		room.Unregister <- client
		ws.Close()
	}()

	// Start goroutines to handle reading and writing messages for this client
	go client.ReadMessages(room)
	go client.WriteMessages()
}
