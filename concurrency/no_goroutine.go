package main

import (
    "fmt"
    "time"
)

func createPizza1(pizza int) {
    time.Sleep(time.Second)
    fmt.Printf("Created pizza %d\n", pizza)
}

func timeTrack1(start time.Time, funcName string) {
    elapsed := time.Since(start)
    fmt.Println(funcName, "took", elapsed)
}

func main() {
    defer timeTrack1(time.Now(), "Pizzaaa...")

    num_of_ovens := 3
    for i := 0; i < num_of_ovens; i++ {
        createPizza1(i)
    }
}
