package operation

import (
	"fmt"
	"github.com/qntfy/kazaam/transform"
	"testProject/learning/example/json_parser/transform/function"
	"testProject/learning/example/json_parser/transform/helper"
)

// custom "shift" operation -> support apply custom func from value "$FUNC(\"old.key.path\")"
//
// extract old.key.path first (using "shift") then apply data with FUNC
func FuncExtract(spec *transform.Config, data []byte) ([]byte, error) {
	var outData []byte
	if spec.InPlace {
		outData = data
	} else {
		outData = []byte(`{}`)
	}
	for k, v := range *spec.Spec {
		array := true
		var keyList []string

		// check if `v` is a string or list and build a list of keys to evaluate
		switch v.(type) {
		case string:
			keyList = append(keyList, v.(string))
			array = false
		case []interface{}:
			for _, vItem := range v.([]interface{}) {
				vItemStr, found := vItem.(string)
				if !found {
					return nil, transform.ParseError(fmt.Sprintf("Warn: Unable to coerce element to json string: %v", vItem))
				}
				keyList = append(keyList, vItemStr)
			}
		default:
			return nil, transform.ParseError(fmt.Sprintf("Warn: Unknown type in message for key: %s", k))
		}

		// iterate over keys to evaluate
		for _, v := range keyList {
			// var vPath = v
			var dataForV []byte
			// var err error

			// check $FUNC(path)
			applyFunc, fnName, fnArgs, vPath, err := helper.IsCustomFunc(v)
			if err != nil {
				return nil, err
			}

			// get data from path
			dataForV, err = GrabData(spec, vPath, data)
			if err != nil {
				return nil, err
			}

			// if array flag set, encapsulate data
			if array {
				// bookend() is destructive to underlying slice, need to copy.
				// extra capacity saves an allocation and copy during bookend.
				tmp := make([]byte, len(dataForV), len(dataForV)+2)
				copy(tmp, dataForV)
				dataForV = helper.Bookend(tmp, '[', ']')
			}

			// apply FUNC before set data back to json
			if applyFunc {
				dataForV, err = function.ApplyFunc(fnName, fnArgs, dataForV)
				if err != nil {
					return nil, err
				}
			}

			// Note: following pattern from current Shift() - if multiple elements are included in an array,
			// they will each successively overwrite each other and only the last element will be included
			// in the transformed data.
			outData, err = helper.SetJSONRaw(outData, dataForV, k, spec.KeySeparator)
			if err != nil {
				return nil, err
			}
		}
	}
	return outData, nil
}

func GrabData(spec *transform.Config, key string, data []byte) ([]byte, error) {
	if key == "$" {
		// ???
		return data, nil
	} else {
		// path
		dataForV, err := helper.GetJSONRaw(data, key, spec.Require, spec.KeySeparator)
		if err != nil {
			return nil, err
		}

		return dataForV, nil
	}
}
