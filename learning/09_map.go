package main

import (
	"fmt"
)

func main()  {
	fmt.Println("key: value")
	m1 := make(map[string]int)

	fmt.Println(m1)
	fmt.Println("- m1 with make() always != nil", m1 != nil)

	var m2 map[string]int
	fmt.Println(m2)
	fmt.Println("- m2 with var always == nil", m2 == nil)

	m3 := map[string]int {
		"k1": 1,
		"k2": 2,
		"k3": 3,
	}
	fmt.Println(m3)

	fmt.Println("- truy van key")
	val, found := m3["k3"]
	fmt.Println("key3 found =", found, "with val =", val)

	val, found = m3["k4"]
	fmt.Println("key4 found =", found, "with val =", val)

	fmt.Println("- them key")
	m3["new_key"] = 4
	m3["Thích"] = 5

	fmt.Println("- loop map")
	for key, val := range m3 {
		fmt.Println(key, ":", val)
	}

	fmt.Println("- delete key")
	delete(m3, "new_key")
	delete(m3, "not_exist_key") // ok
	fmt.Println(m3)

	fmt.Println("- map la reference type")
	m4 := m3
	m4["key_from_m4"] = 6
	fmt.Println("m3 =", m3)
	fmt.Println("m4 =", m4)

	// TODO: map ko compare duoc
}
