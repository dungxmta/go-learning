package function

import (
	"log"
)

// do not encode to json here, just return it as normal data
func Demo(data []byte, params ...interface{}) (out interface{}, err error) {
	log.Println(params, " | ", string(data))

	// out = "ok"
	out = []byte("ok")
	// out = map[string]interface{} {"k": 1, "k2": true, "k3": "v"}
	// out = []interface{} {1, true, "2", map[string]interface{} {"k": 1}}
	// out, err = json.Marshal("ok")
	return
}
