package main

import "net"

type Client struct {
	Id   string
	Room string
	Conn *net.Conn
}

func NewClient(id, room string, conn *net.Conn) *Client {
	return &Client{
		Id:   id,
		Room: room,
		Conn: conn,
	}
}

func (c *Client) Publish(msg string) error {
	// conn := c.Conn
	_, err := (*(c.Conn)).Write([]byte(msg))
	// _, err := (*conn).Write([]byte(msg))
	return err
}
