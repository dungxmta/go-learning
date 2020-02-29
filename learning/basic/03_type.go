package main

import (
    "fmt"
)

func main() {
    // log.Println("Hello go!")
    // int int8,16,32,64
    // uint =/> int +
    // byte =/> alias uint8
    // float32,64
    // complex64,128 =/> z = a + bi
    // bool
    // string
    // rune =/> alias int32
    fmt.Println("\n--- byte")
    var b byte = 255 // 256 will cause err
    fmt.Println(b)
    fmt.Printf("%T \n", b)
    fmt.Printf("%X \n", b)

    fmt.Println("\n---")

    var b2 byte = 'A'
    fmt.Println(b)
    fmt.Printf("%T \n", b2)
    fmt.Printf("%X \n", b2) // hexa

    // fmt.Println("\n--- string")
    // ms := "Tiếng việt"
    // fmt.Println(ms)

    fmt.Println("--- string unicode")
    var r string = "Thích" // 4 kí tự =/> 6 byte
    fmt.Println(r)
    fmt.Printf("len(r)=%d \n", len(r))
    for i := 0; i < len(r); i++ {
        // 6 byte
        fmt.Printf("%c", r[i]) // in ra sẽ ko được kí tự unicode
    }

    fmt.Println("\n--- rune")
    rules := []rune(r) // ep kieu sang rule
    fmt.Println(rules)
    for i := 0; i < len(rules); i++ {
        fmt.Printf("%c", rules[i])
    }
    fmt.Println("\n---")
    var r2 rune = 'A'
    fmt.Printf("%c \n", r2)
}
