package main

import (
    "log"
    "time"
)

/**
Goroutine giao tiếp nhau qua channel
1. Unbuffered channel
2. Buffered channel
*/
func unbufferedChannel() {
    // Unbuffered channel
    ch1 := make(chan int)

    go func() {
        // put 100 to channel
        ch1 <- 100
        // bị block tới khi giá trị được lấy ra
        // nếu ko có code lấy ra thì đoạn code sau sẽ ko được thục thi:
        log.Println("sent")
    }()

    // khi lấy giá trị từ channel cũng bị block
    // nếu ko có thì sẽ raise deadlock
    log.Println(<-ch1)
}

func bufferedChannel() {
    // Buffered channel

    // gửi 2 giá trị vào channel
    // nếu = 0 <-> unbuffered_channel
    ch2 := make(chan int, 2)

    ch2 <- 1
    ch2 <- 2

    close(ch2)
    // ch2 <- 3  // deadlock do capacity=2
    log.Println(<-ch2)
    log.Println(<-ch2)
    log.Println(<-ch2) // deadlock! để ko raise thì cần close channel trước
}

func selectChannel() {
    queue := make(chan int)
    done := make(chan bool)

    // cho phần tử vào queue
    go func() {
        for i := 0; i < 10; i++ {
            log.Println("put i to channel queue: ", i)
            queue <- i
            log.Println("done put queue: ", i)
        }
        done <- true
    }()

    // lấy phần tử từ queue, dừng khi done=true
    for {
        // đợi giá trị lấy ra từ các channel
        select {
        case v := <-queue:
            log.Println(v)
        case <-done:
            log.Println("done")
            return
        default:
            log.Println("...")
            time.Sleep(time.Second)
        }
    }
}

func loopChannel()  {
    queue := make(chan int)

    go func() {
        for i:=0; i<10;i++{
            queue<-i
        }
        close(queue)
    }()

    for val := range queue {
        log.Println(val)
    }
}

func main() {
    // unbufferedChannel()
    // bufferedChannel()
    selectChannel()
    // loopChannel()
    log.Println("done")
}
