package task

import (
	_type "BasicOA/type"
	"net/http"
)
import "github.com/gorilla/websocket"

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	Channel = make(chan _type.Respond)
	clients = make(map[*websocket.Conn]_type.User)
)
