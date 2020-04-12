package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a Message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong Message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum Message size allowed from peer.
	maxMessageSize = 512
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connection is an middleman between the websocket Connection and the Hub.
type Connection struct {
	// The websocket Connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket Connection to the Hub.
func (s Subscription) readPump() {
	c := s.conn

	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))

	c.ws.SetPongHandler(func(string) error {
		err := c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return err
	})

	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := Message{msg, s.room}
		h.broadcast <- m
	}
}

// write writes a Message with the given Message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the Hub to the websocket Connection.
func (s *Subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil)
	vars := mux.Vars(r)

	log.Println(vars["room"])

	if err != nil {
		log.Println(err)
		return
	}
	c := &Connection{send: make(chan []byte, 256), ws: ws}
	s := Subscription{c, vars["room"]}

	h.register <- s

	go s.writePump()

	s.readPump()
}
