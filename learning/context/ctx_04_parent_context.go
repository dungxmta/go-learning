package main

import (
	"context"
	"log"
	"time"
)

// ***NOTE: cancel parent context not cancel its child goroutine
// cancel context not mean return func, only send value to channel context.Done()

// Child context will get Done too when Parent context canceled
func main() {
	log.Println("begin main...")
	parentCtx, cancelParent := context.WithCancel(context.Background())

	childTimeoutCtx, _ := context.WithTimeout(parentCtx, time.Second*10)

	time.AfterFunc(time.Second*2, func() {
		log.Println("Cancel parent context after 2s")
		cancelParent()
	})

	select {
	case <-childTimeoutCtx.Done():
		log.Println("Child context done! Normally this will called after Child ctx timeout (10s)")
	}

	// cancelParent()
	// for i:=0;i<15 ;i++  {
	// 	fmt.Println("...")
	// 	time.Sleep(time.Second)
	// }
}
