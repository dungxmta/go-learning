package main

import (
	"fmt"
	"testProject/learning/example/tcp_chat/config"
)

func main() {
	addr := fmt.Sprintf("%v:%v", config.HOST, config.PORT)

	srv := NewServer(config.PROTOCOL, addr)
	srv.Listen()
}
