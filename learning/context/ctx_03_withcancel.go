package main

import (
	"context"
	"fmt"
	"time"
)

// main end khi nha duoc signal cancel
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(time.Second, func() {
		cancel()
	})

	select {
	case <-ctx.Done():
		fmt.Println("done")
	}
}
