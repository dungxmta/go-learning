package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// https://golangbot.com/read-files/

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

func readSmallChunks(fPath string) error {
	file, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	results := make([]byte, 50)

	for {
		lenReadBytes, err := reader.Read(results)
		if err != nil {
			return err
		}
		fmt.Println(string(results[0:lenReadBytes]))
	}
}

func readFull(fPath string) error {
	// reading entire contents into memory
	// return: []byte
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func seekNPeek(fPath string) error {
	file, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many actually were read.
	b1 := make([]byte, 5)
	n1, _ := file.Read(b1)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// You can also Seek to a known location in the file and Read from there.
	o2, _ := file.Seek(6, 0)
	b2 := make([]byte, 2)
	n2, _ := file.Read(b2)
	fmt.Printf("%d bytes @ %d: %v\n", n2, o2, string(b2[:n2]))

	// same as above, using pkg "io"
	o3, _ := file.Seek(6, 0)
	b3 := make([]byte, 2)
	n3, _ := io.ReadAtLeast(file, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// pkg "bufio"
	r4 := bufio.NewReader(file)
	b4, _ := r4.Peek(5)
	fmt.Printf("5 bytes: %s\n", string(b4))

	return nil
}

func main() {
	for idx, arg := range os.Args {
		fmt.Println(idx, "|", arg)
	}

	// err := readLine("file_1.txt")
	// err := readSmallChunks("file_1.txt")
	// err := readFull("file_1.txt")
	err := seekNPeek("file_1.txt")
	if err != nil {
		log.Fatal(err)
	}
}
