package main

import (
    "fmt"
    "time"
)

func pinger(c chan string)  {
    for i := 0; ; i++ {
        c <- fmt.Sprintf("ping-%d", i)
    }
}

func printer(c chan string)  {
    for {
        msg :=  <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)  // nhap 1 ki tu de ket thuc
    fmt.Println("end main!")
}