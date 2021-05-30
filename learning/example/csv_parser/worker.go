package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

type void struct{}

// func NewDict() *map[string]Row {
// 	return &map[string]Row{}
// }

type Row struct {
	Type  string // ip / cid / sid
	Value string
	Group string
	Owner string
}

/*
{
	value,
	group,
	owner,
	type
}
*/

func startWkPool(maxWk int, files []string, loopCh chan<- Row) *sync.WaitGroup {
	poolCh := make(chan void, maxWk)

	var wg sync.WaitGroup

	for _, file := range files {
		poolCh <- void{} // 1 slot / 1 worker
		wg.Add(1)

		go func(filePath string, poolCh <-chan void, loopCh chan<- Row, wg *sync.WaitGroup) {
			log.Printf("start read %v...", filePath)
			defer log.Printf("%v done", filePath)

			worker(filePath, loopCh)
			<-poolCh // free slot after done
			wg.Done()
		}(file, poolCh, loopCh, &wg)
	}

	return &wg
}

func worker(filePath string, loopCh chan<- Row) {
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

func filter(row []string, loopCh chan<- Row) {

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

	// log.Println("...")
	if gr == "0" {
		gr = ""
	}
	if owner == "0" {
		owner = ""
	}

	loopCh <- Row{
		Type:  typ,
		Value: row[2],
		Group: gr,
		Owner: owner,
	}
}
