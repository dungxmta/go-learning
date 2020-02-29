package main

import "fmt"

func main() {
    /**
        + Array phai khai bao chieu dai co dinh
          Slide ko can
        + arr1 = arr2, khi thay doi arr1 thi arr2 ko doi
          slide thi co
     */
    fmt.Println("=== Array")
    var a1 [2]int
    a1[0] = 1

    var a2 = [2]int {1, 2}
    fmt.Println(a1, a2)

    fmt.Println("- khao bao [...]")
    a3 := [...]string {"1", "2", "3"}
    fmt.Println(a3)

    fmt.Println("- arr la value type, ko phai reference type")
    a4 := a3
    a4[0] = "new"
    fmt.Println(a3)
    fmt.Println(a4)

    fmt.Println("=== Slide")
    /**
        Slices không có bất kì dữ liệu nào. Chúng là các tham chiếu đến mảng hiện có, nó mô tả một phần (hoặc toàn bộ) Array. Nó có kích thước động nên thường được sử dụng nhiều hơn Array.
        Slices có thể tạo ra từ một Array bằng cách lấy từ vị trí phần tử bắt đầu và kết thúc trong Array.
     */
    var s1 []int
    fmt.Println(s1) // []
    // s1[0] = 1 // panic: runtime error: index out of range [0] with length 0

    fmt.Println("- khai bao slide []int {}")
    s2 := []int {1,2,3}
    fmt.Println(s2)

    fmt.Println("- khai bao slide tu Array")
    a5 := [...]int {1,2,3}
    s3 := a5[1:3] // index [1, 2]
    fmt.Println("s3 = ", s3) // [2 3]

    fmt.Println("- khi thay doi slide =/> thay doi toan bo cac slide co ref voi arr ")
    s4 := a5[:]
    a6 := a5
    fmt.Println("s4 = ", s4)

    s4[0] = 0 // s3 ko doi do chi thay doi a[0]
    fmt.Println("s3 = ", s3)

    s4[1] = 1 // s3 thay doi do co ref voi a5[1]
    fmt.Println("s4 = ", s4)
    fmt.Println("a5 = ", a5)
    fmt.Println("s3 = ", s3)
    fmt.Println("a6 = ", a6)

    fmt.Println("- length va capacity ")
    a7 := [...]string {"0", "1", "2", "3", "4", "5"}
    s5 := a7[1:3]

    // len() la so phan tu cua slice
    fmt.Println("len(s5)=", len(s5))

    // cap() la so phan tu tu vi tri start luc ta slice -> cuoi array
    // cap(s5) = len(["1", "2", "3", "4", "5"])
    fmt.Println("cap(s5)=", cap(s5))

    fmt.Println("- make")
    s6 := make([]int, 2, 5)
    fmt.Println(s6, len(s6), cap(s6))

    s6[0] = 0
    s6[1] = 1
    fmt.Println(s6, len(s6), cap(s6))

    fmt.Println("- append")
    s7 := make([]int, 2) // [0 0]
    fmt.Println(s7, len(s7), cap(s7))

    s7 = append(s7, 2, 3) // [0 0 2 3]
    fmt.Println(s7, len(s7), cap(s7))

    s8 := append(s7, 4, 5) // s7 not change
    fmt.Println(s7, len(s7), cap(s7))
    fmt.Println(s8, len(s8), cap(s8))

    //	slice = append(slice, elem1, elem2)
    //	slice = append(slice, anotherSlice...)
    // As a special case, it is legal to append a string to a byte slice, like this:
    //	slice = append([]byte("hello "), "world"...)
    s9 := append(s7, s8...)
    fmt.Println(s9, len(s9), cap(s9))

    fmt.Println("- copy")
    src := []string {"0", "1", "2", "3", "4", "5"}
    dst := make([]string, 2)

    copy_num := copy(dst, src)
    fmt.Println(dst, "copy num = ", copy_num)

    fmt.Println("- delete using append()")
    new_src := append(src[:1], src[2:]...) // remove index = 1
    fmt.Println(new_src)
}
