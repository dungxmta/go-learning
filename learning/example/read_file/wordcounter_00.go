// === test 00: no goroutine  ===
// ===========================================
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type WordsCounter struct {
	found map[string]int
}

func (w *WordsCounter) add(word string) {
	w.found[word] += 1
}

func NewWordsCounter() *WordsCounter {
	return &WordsCounter{found: make(map[string]int)}
}

var files = []string{
	"file_1.txt",
	"file_2.txt",
	"file_3.txt",
	"file_4.txt",
	"file_5.txt",
	"file_6.txt",
	"file_7.txt",
}

func main() {
	wordsCounter := NewWordsCounter()

	// read file each worker
	for index, srcFile := range files {
		name := fmt.Sprintf("Worker_%v", strconv.Itoa(index))

		log.Println("Start ->", name)

		file, err := os.Open(srcFile)
		if err != nil {
			log.Fatal("Error when process data from file", srcFile, err)
		}

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			word := scanner.Text()
			lowerCase := strings.ToLower(word)

			wordsCounter.add(lowerCase)
		}

		log.Println("Done <-", name)
		file.Close()
	}

	log.Println("Done!", len(wordsCounter.found))
	// for k, v := range wordsCounter.found {
	// 	fmt.Printf("%20s : %3v\n", k, v)
	// }
}
