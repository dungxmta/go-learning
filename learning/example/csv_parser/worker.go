package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
	"testProject/learning/example/csv_parser/model"
)

type void struct{}

// func NewDict() *map[string]Row {
// 	return &map[string]Row{}
// }

/*
{
	value,
	group,
	owner,
	type
}
*/

func startWkPool(maxWk int, files []string, loopCh chan<- model.Row, wg *sync.WaitGroup) {
	poolCh := make(chan void, maxWk)

	wg.Add(len(files))

	for _, file := range files {
		poolCh <- void{} // 1 slot / 1 worker

		go func(filePath string, poolCh <-chan void, loopCh chan<- model.Row, wg *sync.WaitGroup) {
			log.Printf("start read %v...", filePath)
			defer log.Printf("%v done", filePath)

			worker(filePath, loopCh)
			<-poolCh // free slot after done
			wg.Done()
		}(file, poolCh, loopCh, wg)
	}
}

func worker(filePath string, loopCh chan<- model.Row) {
	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	reader := csv.NewReader(fi)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		filter(row, loopCh)
	}

}

func filter(row []string, loopCh chan<- model.Row) {

	var (
		typ   = "ip"
		gr    = row[3]
		owner = row[4]
	)

	if isIPv4(row[2]) {
		// check ip

		// TODO: not private =/> skip
		// if !isPrivate(row[2]) {
		// 	return
		// }

	} else {
		// check source
		switch row[1] {
		case "A":
			typ = "cid"
		case "B":
			typ = "sid"
		default:
			// TODO: extract IP?
			return
		}
	}

	// verify data
	if gr == "0" {
		gr = ""
	}
	if owner == "0" {
		owner = ""
	}
	if typ != "ip" && gr == "" && owner == "" {
		return
	}

	loopCh <- model.Row{
		Type:  typ,
		Value: row[2],
		Group: gr,
		Owner: owner,
	}
}
