package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testProject/learning/example/tcp_chat/config"
)

func main() {
	// get nickname & room
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("NAME: ")
	id, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	id = id[:len(id)-1]

	fmt.Print("ROOM: ")
	room, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	room = room[:len(room)]

	log.Println("[debug]", id)
	log.Println("[debug]", room)

	// join room
	client := NewClient(id, room)

	addr := fmt.Sprintf("%v:%v", config.HOST, config.PORT)
	client.Connect(config.PROTOCOL, addr)
	defer (*(client).Conn).Close()

	// chat...
	go client.Type()
	client.Listen()
}
