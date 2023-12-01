package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var updrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade( w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	updrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := updrader.Upgrade(w, r, nil)

	if err != nil {
        log.Println(err)
        return nil, err
    }

	return conn, nil
}