package main

import (
	"context"
	"fmt"
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
	// main end after 5s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel() // tra ve gia tri trong chanel Done

	// WithTimeout_01(ctx)
	// WithTimeout_02(ctx)
	// WithTimeout_03(ctx, cancel)
	WithTimeout_04(ctx, cancel)
	fmt.Println("end main!")
}

// wait 5s
func WithTimeout_01(ctx context.Context) {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
	fmt.Println("end WithTimeout_01()")
}

func WithTimeout_02(ctx context.Context) {
	// ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			fmt.Println("end WithTimeout_02()")
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("...")
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
			fmt.Println(ctx.Err())
			fmt.Println("end WithTimeout_03()")
			return
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("...")
		}
	}
}

func WithTimeout_04(ctx context.Context, cancelFunc context.CancelFunc) {
	canceled := make(chan bool)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			canceled <- true
		}
	}()

	if isCanceled := <-canceled; isCanceled {
		close(canceled)

		defer func() {
			fmt.Println("end WithTimeout_04() after cancel signal")
		}()

		return
	}

	time.Sleep(time.Second * 10)
	fmt.Println("end WithTimeout_04() after 10s")
}
