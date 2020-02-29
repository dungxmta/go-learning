package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test02(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			fmt.Println("after put i to channel: ", i, msg)
		}
	}()

	// go func() {
	//     for true {
	//         fmt.Println("...", <-c)
	//     }
	// }()
	return c
}

func main() {
	c := test02("1st")
	// d := test02("2nd")
	// test02("1st")

	for i := 0; i < 5; i++ {
		fmt.Println("Main: ", <-c)
		// fmt.Println("Main: ", <-d)
	}
	// time.Sleep(time.Second * 5)
	fmt.Println("done...")
}
