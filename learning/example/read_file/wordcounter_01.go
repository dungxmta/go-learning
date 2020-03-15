// === test 01: use channel to communicate ===
// goroutine:   each worker read word from file and put it to words channel
// main:        get value from counter channel then calculate it
// ===========================================
// === >>> TOO SLOW event worse than test_00 with NO goroutine!!!
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

type Job interface {
	Process(chan<- string) error
	Done()
	IsDone() bool
	GetID() string
}

type Worker struct {
	name    string
	srcFile string
	done    chan bool
	// done    context.CancelFunc
}

func (wk *Worker) Process(wordsChannel chan<- string) error {
	file, err := os.Open(wk.srcFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		lowerCase := strings.ToLower(word)

		wordsChannel <- lowerCase
	}

	return scanner.Err()
}

func (wk *Worker) Done() {
	// wk.done <- true
	// wk.done()
	close(wk.done)
}

func (wk *Worker) GetID() string {
	return wk.name
}

func (wk *Worker) IsDone() bool {
	select {
	case <-wk.done:
		return true
	default:
		return false
	}
}

// func NewWorker(name, srcFile string, cancelFunc context.CancelFunc) Job {
func NewWorker(name, srcFile string) *Worker {
	return &Worker{
		name:    name,
		srcFile: srcFile,
		done:    make(chan bool),
	}
}

type WordsCounter struct {
	found map[string]int
}

func (w *WordsCounter) add(word string) {
	w.found[word] += 1
}

func NewWordsCounter() *WordsCounter {
	return &WordsCounter{found: make(map[string]int)}
}

// var wg sync.WaitGroup

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
	wordsChannel := make(chan string)
	// wordsCounter := make(map[string]int)
	wordsCounter := NewWordsCounter()

	// wg.Add(len(files) + 1)
	// wkMap := make(map[string]context.Context)
	wkMap := make(map[string]Job)

	// read file each worker
	for index, srcFile := range files {
		name := fmt.Sprintf("Worker_%v", strconv.Itoa(index))
		// ctx, cancelFunc := context.WithCancel(context.Background())
		// wk := NewWorker(name, srcFile, cancelFunc)
		wk := NewWorker(name, srcFile)

		go func(wk *Worker, srcFile string, wordsChannel chan string) {
			// defer wg.Done()
			defer func() {
				// wk.done = true
				wk.Done()
				log.Println("Done <-", wk.GetID())
			}()

			log.Println("Start ->", wk.GetID())
			err := wk.Process(wordsChannel)
			if err != nil {
				log.Fatal("Error when process data from file", srcFile, err)
			}
		}(wk, srcFile, wordsChannel)

		wkMap[wk.GetID()] = wk
		// wkMap[wk.GetID()] = ctx
	}

	// counter
	// go func(wsc *WordsCounter) {
	for {
		select {
		case word := <-wordsChannel:
			// Counter(word, &wordsCounter)
			// wsc.add(word)
			wordsCounter.add(word)
		default:
			if len(wkMap) == 0 {
				goto ALL_WORKER_DONE
			}
			// fmt.Println("-->before", wkMap)
			for wkID, wk := range wkMap {
				// fmt.Println(wkID, "|", wk.IsDone())
				if wk.IsDone() {
					// fmt.Println("Try delete worker", wkID, "...")
					delete(wkMap, wkID)
					// fmt.Println("-->after", wkMap)
				}
				// TODO: fatal error: all goroutines are asleep - deadlock! ???
				// select {
				// case <-wk.(*Worker).done:
				// 	fmt.Println("Try delete worker", wkID, "...")
				// 	// delete(wkMap, wkID)
				// 	fmt.Println("-->after", wkMap)
				// }
			}
			// time.Sleep(time.Second)
			// fmt.Println("...")
		}
	}
	// }(wordsCounter)
ALL_WORKER_DONE:
	log.Println("Done!", len(wordsCounter.found))
	// for k, v := range wordsCounter.found {
	// 	fmt.Printf("%20s : %3v\n", k, v)
	// }
}
