package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

const (
	dataDir = "./data"
	outDir  = "./output"

	maxWk = 5
)

func main() {
	defer log.Println("DONE!")

	loopCh := make(chan map[string]interface{}, 1000)
	doneCh := make(chan struct{})

	var wg sync.WaitGroup

	go WriteCSV("raw.csv", loopCh, doneCh)

	maxReq := calcTotalRequest(2021, 06, 1, 365, 600)
	log.Println("Max req:", maxReq)
	wg.Add(maxReq)

	go StartPool(maxWk, maxReq, loopCh, &wg)

	wg.Wait()
	close(loopCh)
	log.Println("loopCh closed! wait for WriteCSV done...")
	<-doneCh
}

// http request -> (filter) -> write csv

func StartPool(maxWk, maxReq int, loopCh chan<- map[string]interface{}, wg *sync.WaitGroup) {
	poolCh := make(chan struct{}, maxWk)

	for i := 0; i < maxReq; i++ {
		poolCh <- struct{}{}

		go func(i int) {
			log.Println("start request -", i)
			request(loopCh)
			<-poolCh
			wg.Done()
			log.Println("done request -", i)
		}(i)
	}
}

func calcTotalRequest(yy, mm, dd, duration, interval int) int {
	// t := time.Date(yy, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
	// ts := t.UnixNano()
	return duration * 86400 / interval
}

func request(loopCh chan<- map[string]interface{}) {
	var v string

	for i := 0; i < 100; i++ {
		// time.Sleep(time.Millisecond * 200)
		v = fmt.Sprintf("%v", i)
		loopCh <- map[string]interface{}{
			"id":      v,
			"ip":      v,
			"owner":   v,
			"group":   v,
			"p_group": v,
		}
	}
}

func WriteCSV(filePath string, loopCh <-chan map[string]interface{}, doneCh chan<- struct{}) {
	defer close(doneCh)

	fo, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	writer := csv.NewWriter(fo)
	var line []string

	for row := range loopCh {
		line = extractLine(row)

		err = writer.Write(line)
		if err != nil {
			panic(err)
		}

		line = nil
	}
	writer.Flush()
}

func extractLine(row map[string]interface{}) (line []string) {
	id, _ := row["id"].(string)
	ip, _ := row["ip"].(string)
	owner, _ := row["owner"].(string)
	gr, _ := row["group"].(string)
	pgr, _ := row["p_group"].(string)

	line = append(line, id, ip, owner, gr, pgr)
	return
}
