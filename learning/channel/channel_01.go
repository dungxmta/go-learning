package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		fmt.Println("after put i to channel: ", i)
	}
}

func main() {
	c := make(chan string)
	go test("...", c)

	for i := 0; i < 5; i++ {
		fmt.Println("Main: ", <-c)
	}
	fmt.Println("done...")
}
