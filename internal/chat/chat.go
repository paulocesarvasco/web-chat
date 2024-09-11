package chat

import "log"

type ChatRoom struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (room *ChatRoom) Run() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
		case client := <-room.Unregister:
			if _, ok := room.Clients[client]; ok {
				delete(room.Clients, client)
				close(client.Send)
				log.Printf("client disconnected: %s\n", client.Conn.LocalAddr().String())
			}
		case message := <-room.Broadcast:
			for client := range room.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(room.Clients, client)
					log.Printf("client disconnected: %s\n", client.Conn.LocalAddr().String())
				}
			}
		}
	}
}
