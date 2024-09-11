package ws

import (
	"github.com/gorilla/websocket"
	// "log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
