package main

import (
    "fmt"
)

func applyPointer(p *string)  {
	*p += "_from_applyPointer"
}

func main() {
    ms := "Tiếng việt"

    var p1 *string // value = nil
    p1 = &ms

    fmt.Println(ms, &ms)
    fmt.Println(p1, *p1, &p1)

    p2 := new(string) // value != nil
    fmt.Println(p1)
    fmt.Println(p2)

    *p1 = *p1 + "_from_p2"
    fmt.Println(ms)

    fmt.Println("pointer -> array")
    arr := [3]int{1, 2, 3}
    p4 := &arr

    var p5 *[3]int
	p5 = &arr

	fmt.Println(p4, *p4, &p4)
	fmt.Println(p5, *p5, &p5)

	fmt.Println("func(pointer)")
	applyPointer(p1)
	fmt.Println(ms)
}
