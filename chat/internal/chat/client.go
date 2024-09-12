package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func (client *Client) ReadMessages(room *ChatRoom) {
	defer func() {
		room.Unregister <- client
		client.Conn.Close() // Ensure the connection is properly closed
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			// Log and handle the error
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		// Broadcast the message to the chat room
		room.Broadcast <- message
	}
}

func (client *Client) WriteMessages() {
	defer client.Conn.Close() // Ensure the connection is properly closed

	for message := range client.Send {
		if err := client.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}
