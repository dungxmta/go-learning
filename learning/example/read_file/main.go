package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLine(fPath string) error {
	file, err := os.Open(fPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	return scanner.Err()
}

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, "|", arg)
	}

	err := readLine("file_1.txt")
	if err != nil {
		log.Fatal(err)
	}
}
