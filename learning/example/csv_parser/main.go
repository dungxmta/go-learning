package main

import (
	"log"
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

	loopCh := make(chan Row, 100)
	doneCh := make(chan void)
	dict := &map[string]Row{}

	wg := startWkPool(maxWk, files, loopCh)
	go normalizer(dict, loopCh, doneCh)

	wg.Wait()
	close(loopCh)
	log.Println("loopCh closed! wait for producer done...")
	<-doneCh

	pr := NewProducer()
	pr.writeCSV(dict)
}

func normalizer(dict *map[string]Row, loopCh <-chan Row, doneCh chan<- void) {
	defer close(doneCh)

	// var (
	// 	val, gr, owner, typ string
	// )
	var (
		added bool
		old   Row
	)

	for row := range loopCh {
		// log.Println(row)

		if old, added = (*dict)[row.Value]; added {
			// only update if not ok
			if old.Owner == "" && row.Owner != "" {
				old.Owner = row.Owner
			}
			if old.Group == "" && row.Group != "" {
				old.Group = row.Group
			}
			continue
		}

		(*dict)[row.Value] = row
	}

	// log.Println(dict)
}
