package main

import "fmt"

func main() {
    switch num := 10; num {
    case 1, 2, 3, 4:
        fmt.Println("wrong")
    case 10:
        fmt.Println("10")
        fallthrough // tiep tuc
    case 5:
        fmt.Println("fallthrough - 5")
        // fallthrough
    default:
        fmt.Print("default")
    }

    // labelName:
    //	fmt.Println("goto labelName")

    switch num := 10; num {
    case 1, 2, 3, 4:
        fmt.Println("wrong")
    case 10:
        fmt.Println("10")
        goto labelName
    default:
        fmt.Print("default")
    }

    fmt.Println("ignore me 'cause goto")
labelName:
    fmt.Println("goto labelName")
}
