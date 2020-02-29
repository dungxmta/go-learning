package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
2 channel --go to--> 1 channel in main
*/

func test02(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Println("after put i to channel: ", i, msg)
		}
	}()

	return c
}

func test03(inp1, inp2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-inp1
		}
	}()

	go func() {
		for {
			c <- <-inp2
		}
	}()

	return c
}

func test03_select(inp1, inp2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for true {
			select {
			case v := <-inp1:
				c <- v
			case v := <-inp2:
				c <- v
			case <-time.After(time.Second * 1):
				fmt.Println("...")
				// default:
				//     fmt.Println("...")
				//     time.Sleep(time.Second*1)
			}
		}
	}()
	return c
}

func main() {
	c := test02("1st")
	d := test02("2nd")

	// e := test03(c, d)
	e := test03_select(c, d)

	for i := 0; i < 10; i++ {
		fmt.Println("Main: ", <-e)
	}

	fmt.Println("done...")
}
