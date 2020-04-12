package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single chatting user.
type client struct {
	ID string

	// socket is the web socket for this client.
	socket *websocket.Conn

	// send is a channel on which messages are sent.
	send chan []byte

	// room is the room this client is chatting in.
	room *room
}

// get msg from socket
//
// when client.js do socket.send(...) -> msg will be process here
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

// send data from socket to client
//
// client.js get this msg through socket.onmessage(...)
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
