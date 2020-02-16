package main

import (
    "log"
    "sync"
)

var wg sync.WaitGroup

func task_01() {
    log.Println("task 01")
    // nếu ko dùng Done sẽ raise deadlock: wg ko biết lúc nào kết thúc
    wg.Done()
}

func task_02() {
	defer wg.Done()
    log.Println("task 02")
}

/**
	synchronized goroutines
	- thứ tự chạy các goroutines do goruntime quản lý
	- logic wg: chờ cho các goroutines hoàn thành xong thì mới kết thúc main
 */
func main() {
    log.Println("begin")

    wg.Add(2)
    // neu ko dung wait group -> 2 task co the chua tra ve ket qua nhung main() da ket thuc
    go task_01()
    log.Println("before call task02")
    go task_02()

    wg.Wait()

    log.Println("end")
}
