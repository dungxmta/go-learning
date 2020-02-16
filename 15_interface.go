package main

import "fmt"

type Action interface {
    speak(somethings string)
}
type Action2 interface {
    speak2(somethings string)
}

type EmbedInterface interface {
    Action
    Action2
}

type Person3 struct {
    id   int
    name string
}

func (p Person3) speak(somethings string) {
    fmt.Printf("speaking... %s\n", somethings)
}
func (p Person3) speak2(somethings string) {
    fmt.Printf("speaking2... %s\n", somethings)
}

// empty interface
func funcWithParamInterface(empty_ie interface{}) {
    fmt.Println(empty_ie)
}

func main() {
    // named
    p1 := Person3{
        id:   1,
        name: "A",
    }
    p1.speak("A")

    var a Action = p1
    a.speak("B")

    var ei EmbedInterface = p1
    ei.speak("1")
    ei.speak2("2")

    funcWithParamInterface(1)
    funcWithParamInterface(1.1)
    funcWithParamInterface("A")
    funcWithParamInterface(
        struct {
            int
            string
        }{1, "A"})
}
