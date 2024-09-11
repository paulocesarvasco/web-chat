package chat

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func (client *Client) ReadMessages(room *ChatRoom) {
	defer func() {
		room.Unregister <- client
		client.Conn.Close()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
		room.BroadcastMessage(message)
	}
}

func (client *Client) WriteMessages() {
	defer client.Conn.Close()
	for message := range client.Send {
		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
