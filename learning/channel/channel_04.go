package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
   Each msg timeout after 1s
or
    Whole main will end after 5s
*/

func test02(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			fmt.Println("after put i to channel: ", i, msg)
		}
	}()
	return c
}

func timeOutEachMsg(c <-chan string) {
	quitLoop := false

	for {
		select {
		case tmp := <-c:
			fmt.Println("Main: ", tmp)
		case <-time.After(1 * time.Second):
			fmt.Println("Longer than 1s. Quitting...")
			quitLoop = true
		}
		if quitLoop {
			break
		}
	}
	fmt.Println("End timeOutEachMsg")
}

func timeOutMainAfter(c <-chan string) {
	quitLoop := false
	timeout := time.After(5 * time.Second)
	for {
		select {
		case tmp := <-c:
			fmt.Println("Main: ", tmp)
		case <-timeout:
			fmt.Println("Timeout 5s. Quitting...")
			quitLoop = true
		}
		if quitLoop {
			break
		}
	}
	fmt.Println("End timeOutMainAfter")
}

func main() {
	c := test02("1st")

	// timeOutEachMsg(c)
	timeOutMainAfter(c)

	fmt.Println("done...")
}
