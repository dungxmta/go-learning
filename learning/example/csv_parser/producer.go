package main

import (
	"encoding/csv"
	"os"
	"path"
	"testProject/learning/example/csv_parser/model"
)

type Producer struct{}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) writeCSV(dict *map[string]model.Row) {
	outPath := path.Join(outDir, "data.csv")
	fo, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	writer := csv.NewWriter(fo)
	var line []string

	for _, row := range *dict {
		line = []string{
			row.Type,
			row.Value,
			row.Group,
			row.Owner,
		}
		err = writer.Write(line)
		if err != nil {
			panic(err)
		}
		line = nil
	}
	writer.Flush()
}
