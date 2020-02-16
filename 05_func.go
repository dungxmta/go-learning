package main

import "fmt"

// ================================= normal function
func A() {
    fmt.Println("A")
}

func multipleReturnValues(a, b int) (int, int, int) {
    return a, b, 3
}

func nameReturnValues(a, b int) (na, nb int, nc bool) {
    na = a
    nb = b
    nc = true
    return
    // return a, b, true
}

// ================================= variadic function: ko gioi han truyen tham so
func addItem(item int, list ...int) {
    // 2, 3 -> new slice []int {2, 3}
    list = append(list, item)
    fmt.Println(list)
}

func changeRef(list ...int)  {
    list[0] = 999
}


func main() {
    A()
    fmt.Println(multipleReturnValues(1, 2))
    fmt.Println(nameReturnValues(1, 2))

    fmt.Println("- truyen bien vao variadic func")
    addItem(0, 2, 3)

    fmt.Println("- truyen slice vao variadic func")
    list := []int {2, 3}
    addItem(0, list...)

    fmt.Println("- truyen tham chieu (pass a ref)")
    changeRef(list...)
    fmt.Println(list)
}
