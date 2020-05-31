package function

import "encoding/json"

func ApplyFunc(name string, args []interface{}, data []byte) ([]byte, error) {
	// params, err := helper.ParseParams(args)
	// if err != nil {
	// 	return nil, err
	// }

	var out interface{}
	var err error

	switch name {
	case "DEMO":
		out, err = Demo(data, args...)
	}

	var bbytes []byte

	switch v := out.(type) {
	case []byte:
		// do nothing
		// bbytes = v
		out = string(v)
		// case string:
		// 	bbytes = []byte(v)
		// default:
		// 	// json encode
		// 	bbytes, err = json.Marshal(out)
	}

	// always encode back to json
	bbytes, err = json.Marshal(out)

	return bbytes, err
}
