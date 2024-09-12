package ws

import (
	"github.com/paulocesarvasco/web-chat/chat/internal/chat"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request, room *chat.ChatRoom) {
	req, _ := http.NewRequest(http.MethodGet, "http://auth_service:8081/auth", nil)
	req.SetBasicAuth("user_test", "user_pass")
	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Print("error request auth: ", err)
	} else {
		log.Print(res.StatusCode)
	}

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
