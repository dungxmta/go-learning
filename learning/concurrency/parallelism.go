package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Printf("Created pizza %d\n", pizza)
}

func timeTrack(start time.Time, funcName string) {
	elapsed := time.Since(start)
	fmt.Println(funcName, "took", elapsed)
}

func main() {
	// defer -> put lời gọi func và list, sau khi hàm bao quanh thực thi xong thì gọi lời thực thi được lưu ra
	// vd: khi chạy xong main() -> timeTrack()
	defer timeTrack(time.Now(), "Pizzaaa...")

	numOfOvens := 3
	runtime.GOMAXPROCS(numOfOvens) // 3 process
	wg.Add(numOfOvens)

	for i := 0; i < numOfOvens; i++ {
		go createPizza(i)
	}
	wg.Wait()
}
