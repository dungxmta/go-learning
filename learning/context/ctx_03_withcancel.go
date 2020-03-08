package main

import (
	"context"
	"log"
	"time"
)

// main end khi nhan duoc signal cancel
func main() {
	log.Println("begin main...")
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select {
	case <-ctx.Done():
		log.Println("done")
	}
}
