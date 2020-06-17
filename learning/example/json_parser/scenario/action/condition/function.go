package condition

import (
	"fmt"
	"reflect"
)

// ==== Apply function ====

type fnType func(valueIn interface{}) (out interface{}, err error)

var fnMap = map[string]fnType{
	FnCount:  Count,
	FnStrLen: StrLen,
}

func ApplyFunc(name string, valueIn interface{}) (out interface{}, err error) {
	fn, ok := fnMap[name]
	if !ok {
		return nil, fmt.Errorf("invalid func %v", name)
	}

	out, err = fn(valueIn)
	return
}

// @input type   -   output
//  list            length of list
//  null            0
//  others          1
func Count(valueIn interface{}) (out interface{}, err error) {

	// switch v := valueIn.(type) {
	// case []interface{}: // this may not work???
	// 	out  = len(v)
	// default:
	// 	out = 1
	// }

	if valueIn == nil {
		return 0, nil
	}

	val := reflect.ValueOf(valueIn)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		out = val.Len()
	default:
		out = 1
	}

	return
}

// @input type   -   output
//  string          length of string
//  null            0
//  others          <error>
func StrLen(valueIn interface{}) (out interface{}, err error) {
	if valueIn == nil {
		return 0, nil
	}

	switch v := valueIn.(type) {
	case string:
		out = len(v)
	default:
		out = 0
		err = fmt.Errorf("func StrLen only support value string -> %T given", v)
	}

	return
}
