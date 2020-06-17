package condition

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	FnNone   = "NONE"
	FnCount  = "COUNT"
	FnStrLen = "STR_LEN"
)

const (
	OpEq    = "="
	OpDiff  = "!="
	OpGt    = ">"
	OpLt    = "<"
	OpIn    = "IN"
	OpNotIn = "NOT_IN"
)

func CheckCond() (result bool, err error) {

	paramMap := new(map[string]interface{})

	// 1. get data left
	left, err := GetValue("param_left", "key_left", "function_left", paramMap)
	if err != nil {
		return
	}

	// 2. get data right
	right, err := GetValue("param_right", "key_right", "function_right", paramMap)
	if err != nil {
		return
	}

	// 3. start compare
	result, err = Compare(left, OpEq, right)
	return
}

// ======= helper =======

func GetValue(param, key, fnName string, paramMap *map[string]interface{}) (out interface{}, err error) {
	var valueIn interface{}

	// param? -> get from paramMap
	if param != "" {
		if paramMap == nil {
			return nil, fmt.Errorf("param '%v' not found", param)
		}

		doc, ok := (*paramMap)[param]
		if !ok {
			return nil, fmt.Errorf("param '%v' not found", param)
		}

		// key? -> get from param only if data is MAP, else return error
		switch d := doc.(type) {
		case map[string]interface{}:
			if key != "" {
				// extract
				if strings.Contains(key, "[") || strings.Contains(key, "]") ||
					strings.Contains(key, ".") || strings.Contains(key, "*") {

					valueIn, err = Extract(d, key)
					if err != nil {
						return nil, err
					}

					break
				}
			}

			valueIn = d

		// case []map[string]interface{}:
		// 	if key != "" {
		// 		return nil, fmt.Errorf("cannot config key '%v' of param '%v' in array type", key, param)
		// 	}
		// 	value = d

		default:
			if key != "" {
				return nil, fmt.Errorf("cannot config key '%v' of param '%v' in %T type", key, param, d)
			}

			valueIn = d
		}

	} else { // key become a value
		valueIn = key
	}

	// no func
	if fnName == FnNone || fnName == "" {
		return valueIn, nil
	}

	out, err = ApplyFunc(fnName, valueIn)
	return
}

func Extract(input interface{}, key string) (out interface{}, err error) {
	b, _ := json.Marshal(input)

	bb, err := Transform(b)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bb, &out)
	return
}

func Transform(data []byte) ([]byte, error) {
	// TODO: ...
	return nil, nil
}
