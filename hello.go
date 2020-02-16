package main

import (
	"fmt"
	"log"
)

func main()  {
	log.Println("Hello go!")

	ms := "Tiếng việt"

	for i := range ms {
		fmt.Println("%c", i)
	}
}
