package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Client struct {
	Id   string
	Room string

	Conn *net.Conn
}

func NewClient(id, room string) *Client {
	return &Client{
		Id:   id,
		Room: room,
		Conn: nil,
	}
}

func (c *Client) Connect(proto, addr string) {
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	c.Conn = &conn

	// join room
	joinMsg := fmt.Sprintf("%v:%v", c.Id, c.Room)
	err = c.Chat(joinMsg)
	if err != nil {
		log.Fatal(err)
	}
}

// read from server
func (c *Client) Listen() {
	for {
		reader := bufio.NewReader(*c.Conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			// log.Println("Err: ", err)
			time.Sleep(time.Second * 1)
			// if (*c.Conn)
			continue
		}
		log.Print(msg)
	}
}

// read from keyboard
func (c *Client) Type() {
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		// if err != nil {
		// 	log.Println("[Error] Cant get  ", err)
		// 	break
		// }
		// log.Print(msg)
		err := c.Chat(msg)
		if err != nil {
			log.Fatal("[Error] Cant send message: ", err)
		}
	}
}

func (c *Client) Chat(msg string) error {
	_, err := (*(c.Conn)).Write([]byte(msg))
	return err
}
