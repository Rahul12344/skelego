package skelego

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingTime       = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

//Websocket Interface to handle basic websockets
type Websocket interface {
	WriteToSocket(messageType int, buff []byte) error
	ReadFromSocket() (int, []byte, error)
	CloseSocket() error
}

//socket directly interfaces with websocket
type socket struct {
	conn           *websocket.Conn
	messageChannel chan []byte
}

//New Creates new interface with websocket on upgraded endpoint
func New(w http.ResponseWriter, r *http.Request) (Websocket, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return &socket{
			conn:           ws,
			messageChannel: make(chan []byte),
		}, err
	}
	return &socket{
		conn:           ws,
		messageChannel: make(chan []byte),
	}, nil
}

// WriteToSocket writes a message
func (ws *socket) WriteToSocket(messageType int, buff []byte) error {
	ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return ws.conn.WriteMessage(messageType, buff)
}

// ReadFromSocket reads a message
func (ws *socket) ReadFromSocket() (int, []byte, error) {
	return ws.conn.ReadMessage()
}

// CloseSocket closes connections
func (ws *socket) CloseSocket() error {
	return ws.conn.Close()
}
