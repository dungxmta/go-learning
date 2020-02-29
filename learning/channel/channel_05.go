package main

import (
	"fmt"
	"time"
)

/**
receive on quit channel

   The goroutine end when receive quit signal from main
*/

func cleanup() {
	fmt.Println("Goroutine: clean...")
	time.Sleep(time.Millisecond * 1000)
}

func test05(msg string, quit chan string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
				time.Sleep(time.Millisecond * 500)
				fmt.Println("...", i)
			case mainSayQuit := <-quit:
				fmt.Println(mainSayQuit)
				cleanup()
				quit <- "Goroutine: See ya!"
				return
			}
		}
	}()

	return c
}

func main() {
	quitSignal := make(chan string)
	c := test05("1st", quitSignal)

	for i := 0; i < 10; i++ {
		fmt.Println("Main: ", <-c)
	}
	quitSignal <- "Main: I wanna quit...!"
	fmt.Println(<-quitSignal)

	fmt.Println("done...")
}
