package main

import (
	"context"
	"log"
	"time"
)

/**
type Context interface {
    Done() <- chan struct{}
    Err() error
    Deadline(deadline time.Time, ok bool)
    Value(key interface{}) interface{}
}
*/

func main() {
	log.Println("begin main...")
	// main end after 5s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel() // tra ve gia tri trong chanel Done

	// WithTimeout_01(ctx)
	// WithTimeout_02(ctx)
	// WithTimeout_03(ctx, cancel)
	WithTimeout_04(ctx, cancel)
	log.Println("end main!")
}

// wait 5s
func WithTimeout_01(ctx context.Context) {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
	log.Println("end WithTimeout_01()")
}

func WithTimeout_02(ctx context.Context) {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			log.Println("end WithTimeout_02()")
			return
		default:
			time.Sleep(time.Second * 1)
			log.Println("...")
		}
	}
}

// cancel function when get signal from main
// wait 2s then cancel()
func WithTimeout_03(ctx context.Context, cancelFunc context.CancelFunc) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	time.AfterFunc(time.Second*2, func() {
		cancelFunc()
	})

	// defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			log.Println("end WithTimeout_03()")
			return
		default:
			time.Sleep(time.Second * 1)
			log.Println("...")
		}
	}
}

func WithTimeout_04(ctx context.Context, cancelFunc context.CancelFunc) {
	canceled := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			canceled <- true
		}
	}()

	if isCanceled := <-canceled; isCanceled {
		close(canceled)

		defer func() {
			log.Println("end WithTimeout_04() after cancel signal")
		}()

		return
	}

	time.Sleep(time.Second * 10)
	log.Println("end WithTimeout_04() after 10s")
}
