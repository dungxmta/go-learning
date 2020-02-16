package main

import "fmt"

type Human struct {
    id   int
    name string
}
/**
Defind method

func (t Type) methodName(params) returns { ... }

(t Type) -> Receiver
1. Value Receiver
2. Pointer Receiver
 */
func (h Human) getName() string {
    return h.name
}

// Value Receiver -> ko lam thay doi gia tri
func (h Human) changeName() {
    fmt.Printf("h = %p\n", &h)
    h.name = "new name"
}

// Pointer Receiver -> thay doi gia tri
func (h *Human) changeName2() {
    fmt.Printf("h = %p\n", h)
    h.name = "new_name"
}

// non-struct
type String string
// func (s string) append(str string) -> ko lam duoc voi build-in
func (s String) append(str string) string {
    return fmt.Sprintf("%s, %s", s, str)
}


func main() {
    h1 := Human{
        id:   1,
        name: "A",
    }

    // ko thay doi gia tri
    fmt.Printf("h1 = %p\n", &h1)
    h1.changeName()
    fmt.Println(h1)

    // thay doi gia tri
    h1.changeName2()
    fmt.Println(h1)

    // non-struct
    s := String("S")
    newS := s.append("append")
    fmt.Println(newS)
}
