// === test 02: use shared variable between goroutine  ===
// ===========================================
// === >>> best: using sync.Mutex to lock share variable between goroutine
// ===========================================
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type WordsCounter struct {
	sync.Mutex
	found map[string]int
}

func (w *WordsCounter) add(word string) {
	w.Lock()
	defer w.Unlock()

	w.found[word] += 1
}

func NewWordsCounter() *WordsCounter {
	return &WordsCounter{found: make(map[string]int)}
}

func Process(srcFile string, counter *WordsCounter) error {
	file, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		lowerCase := strings.ToLower(word)

		counter.add(lowerCase)
	}

	return scanner.Err()
}

var wg sync.WaitGroup
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
	counter := NewWordsCounter()

	// read file each worker
	for index, srcFile := range files {
		wg.Add(1)

		name := fmt.Sprintf("Worker_%v", strconv.Itoa(index))

		go func(name string, srcFile string, counter *WordsCounter) {
			defer func() {
				wg.Done()
				log.Println("Done <-", name)
			}()

			log.Println("Start ->", name)
			err := Process(srcFile, counter)
			if err != nil {
				log.Fatal("Error when process data from file", srcFile, err)
			}
		}(name, srcFile, counter)

	}

	log.Println("Wait...")
	wg.Wait()
	log.Println("Done!", len(counter.found))
	// for k, v := range counter.found {
	// 	fmt.Printf("%20s : %3v\n", k, v)
	// }
}
