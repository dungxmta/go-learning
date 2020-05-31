package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	PATTERN = `^\$(?P<func>\w+)\((?P<params>.*?)\)$`
)

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(PATTERN)
}

func IsCustomFunc(v string) (status bool, fnName string, fnArgs []interface{}, vPath string, err error) {
	match := re.FindStringSubmatch(v) // [$TEST(1,2,3) TEST 1,2,3]
	// names := re.SubexpNames() // [ func params]
	if len(match) < 3 {
		return false, "", nil, v, nil
	}

	status = true
	fnName = match[1]

	// default the first args always key path to extract data
	args := match[2] // this contains all args, include vPath

	params, err := ParseParams(args)
	if err != nil {
		return false, "", nil, "", err
	}

	argsNum := len(params)
	if argsNum == 0 {
		return false, "", nil, "", errors.New("missing first args")
	}

	vPath, ok := params[0].(string)
	if !ok {
		return false, "", nil, "", errors.New("invalid first args")
	}

	// skip first param
	if argsNum > 1 {
		fnArgs = params[1:]
	}

	return
}

func ParseParams(v string) ([]interface{}, error) {
	var params []interface{}

	value := strings.Trim(v, "")
	if value == "" {
		return params, nil
	}

	value = fmt.Sprintf("[%v]", value)
	err := json.Unmarshal([]byte(value), &params)

	return params, err
}
