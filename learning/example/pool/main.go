package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func lazyTask() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		log.Println("...", i)
	}
}

func main() {
	defer log.Println("END")
	log.Println("INIT")

	poolSize := 5
	pool := New(poolSize)

	fmt.Println("Enter q/quit to exit program")
	inp := bufio.NewScanner(os.Stdin)

	for i := 0; ; i++ {
		inp.Scan()

		switch inp.Text() {
		case "q", "quit":
			return
		case "n": // fire new task by command
			fmt.Println("> send task to work channel", i)
			pool.work <- lazyTask
			fmt.Println("> done send task to work channel", i)
			break
		default:
			pool.Schedule(lazyTask, i)
		}
	}
}
