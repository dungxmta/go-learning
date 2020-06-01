package function

import (
	"encoding/json"
	"log"
)

// do not encode to json here, just return it as normal data
//
// support ignore empty string or null value in array
// *note: now only support []string
func SkipEmpty(data []byte) (out interface{}, err error) {
	// decode to list[]
	var arr []interface{}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return data, err
	}
	log.Println("--> ", arr)

	var outArr []interface{}

	// ignore empty string or nil value
	for _, item := range arr {
		switch v := item.(type) {
		case string:
			if v == "" {
				log.Println("-> v is empty string")
			} else {
				outArr = append(outArr, v)
			}
		default:
			if v == nil {
				log.Println("-> v is null")
			} else {
				outArr = append(outArr, v)
			}
		}
	}

	return outArr, nil
}
