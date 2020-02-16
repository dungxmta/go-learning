package main

import "fmt"

type Person struct {
    id   int
    name string
}
type Person2 struct {
    id   int
    name string
}

type StudentInfo struct {
    className string
    info      Person
}

func main() {
    // named
    p1 := Person{
        id:   1,
        name: "A",
    }
    p2 := Person{2, "B"}

    fmt.Println(p1.id, p1.name, p1)
    fmt.Println(p2.id, p2.name, p2)

    var p3 Person = struct {
        id   int
        name string
    }{
        id:   3,
        name: "C",
    }
    fmt.Println(p3)

    // anonymous struct
    p4 := struct {
        id   int
        name string
    }{
        id:   4,
        name: "B",
    }
    fmt.Println(p4)

    // pointer -> struct
    pointer := &Person{
        5,
        "D",
    }
    fmt.Println(pointer, &pointer, )
    fmt.Println(pointer.id, pointer.name)
    fmt.Println((*pointer).id, (*pointer).name) // bo qua * cung ok

    // anonymous fields
    type NoName struct {
        int
        string
    }

    nn1 := NoName{1, "A"}
    fmt.Println(nn1, nn1.int, nn1.string)

    // nested struct
    s1 := StudentInfo{
        className: "Math",
        info: Person{
            id:   1,
            name: "A",
        },
    }
    fmt.Println(s1, s1.info.id)

    // TODO: struct compare == duoc khi toan bo cac field compare duoc
    p5 := Person{1, "A"}
    p6 := Person{1, "A"}
    // p7 := Person2{1, "A"}
    fmt.Println(p5 == p6)
    // fmt.Println(p5 == p7) // -> mismatched types
}
