package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Angular json editor: https://www.npmjs.com/package/jsoneditor

const (
	FUNC_PATTERN = `^\$(?P<func>\w+)\((?P<params>.*?)\)$`
)

func main() {
	jsonStr := `{
	"params_str_1": "$TEST()",
    "params_str_2": "$TEST(1,2,3)",
	"params_invalid": "$TEST",

    "params_obj": {
        "syncTime": "$TEST(1,\"1\", [1, \"1\", true], {\"k\":1, \"k2\":[1, \"1\", true]})"
    },
    "params_arr": [
        "$TEST(1,\"1\", [1, \"1\", true], {\"k\":1, \"k2\":[1, \"1\", true]})"
    ],

    "normal_key_1": "just string abc...",
    "normal_key_2": 1,
    "normal_key_3": [1, 2],
    "normal_key_4": null
}`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	log.Println("err -> ", err)
	log.Println("after parse json -> ", data)

	re, err := regexp.Compile(FUNC_PATTERN)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, v)
		switch v.(type) {
		case []interface{}:
			log.Println("slice...")
		case map[string]interface{}:

		case string:
			match := re.FindStringSubmatch(v.(string)) // [$TEST(1,2,3) TEST 1,2,3]
			// names := re.SubexpNames() // [ func params]
			if len(match) < 3 {
				log.Println("-> invalid")
				continue
			}

			funcName := match[1]
			paramStr := match[2]

			log.Printf("func=%v | params=%v", funcName, paramStr)
			params, err := parseParams(paramStr)
			log.Printf("-> params parsed: %v | err: %v\n ", params, err)
			// for i, name := range names  {
			// 	if i > 0 && i <= len(match) {
			// 		funcMapStr[name] = match[i]
			// 	}
			// }
		}
	}

	// var d []interface{}
	// err = json.Unmarshal([]byte(`[1,"2", null]`), &d)
	d, err := parseParams(`1,"2", null`)
	log.Println("err -> ", err)
	log.Println("after parse json -> ", d)
}

func parseParams(v string) ([]interface{}, error) {
	var params []interface{}

	value := strings.Trim(v, "")
	if value == "" {
		return params, nil
	}

	value = fmt.Sprintf("[%v]", value)
	err := json.Unmarshal([]byte(value), &params)

	return params, err
}
