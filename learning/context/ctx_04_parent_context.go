package main

import (
	"context"
	"fmt"
	"time"
)

// Child context will get Done too when Parent context canceled
func main() {
	parentCtx, cancelParent := context.WithCancel(context.Background())

	childTimeoutCtx, _ := context.WithTimeout(parentCtx, time.Second*10)

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("Cancel parent context after 2s")
		cancelParent()
	})

	select {
	case <-childTimeoutCtx.Done():
		fmt.Println("Child context done! Normally this will called after Child ctx timeout (10s)")
	}
}
