package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

// convert []byte{1, 2, 3} -> string [1,2,3]
func GoByte2JsonArray(b []byte) string {
	var tmp []string
	for _, v := range b {
		// fmt.Printf("%v ", string(v))
		tmp = append(tmp, fmt.Sprintf("%v", v))
	}

	s := fmt.Sprintf("[%v]", strings.Join(tmp, ","))
	return s
}

// convert string [1,2,3] -> []byte{1, 2, 3}
func JsonArray2GoByte(s string) ([]byte, error) {
	var b []byte
	err := json.Unmarshal([]byte(s), &b)
	return b, err
}
