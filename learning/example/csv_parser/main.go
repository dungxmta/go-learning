package main

import (
	"log"
	"sync"
	"testProject/learning/example/csv_parser/model"
)

const (
	dataDir = "./data"
	outDir  = "./output"

	maxWk = 5
)

func main() {
	defer log.Println("DONE!")

	files, err := getFilesFromFolder(dataDir)
	if err != nil {
		panic(err)
	}

	log.Println(files)

	if len(files) == 0 {
		log.Println("No file...")
		return
	}

	loopCh := make(chan model.Row, 100)
	doneCh := make(chan void)
	dict := &map[string]model.Row{}

	var wg sync.WaitGroup

	go normalizer(dict, loopCh, doneCh)
	go startWkPool(maxWk, files, loopCh, &wg)

	wg.Wait()
	close(loopCh)
	log.Println("loopCh closed! wait for producer done...")
	<-doneCh

	pr := NewProducer()
	pr.writeCSV(dict)
}

func normalizer(dict *map[string]model.Row, loopCh <-chan model.Row, doneCh chan<- void) {
	defer close(doneCh)

	// var (
	// 	val, gr, owner, typ string
	// )
	var (
		added     bool
		old       model.Row
		rowChange bool
	)

	for row := range loopCh {
		// log.Println(row)
		rowChange = false

		if old, added = (*dict)[row.Value]; added {
			if old.Owner == "" && row.Owner != "" {
				old.Owner = row.Owner
				rowChange = true
			}
			if old.Group == "" && row.Group != "" {
				old.Group = row.Group
				rowChange = true
			}

			if rowChange {
				(*dict)[row.Value] = old
			}
			continue
		}

		(*dict)[row.Value] = row
	}

	// log.Println(dict)
}
