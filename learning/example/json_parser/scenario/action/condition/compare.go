package condition

import "fmt"

type fnOperatorType func(left interface{}, right interface{}) (result bool, err error)

var operatorMap = map[string]fnOperatorType{
	OpEq:   OperatorEqual,
	OpDiff: OperatorDiff,
	// OpGt    = ">"
	// OpLt    = "<"
	// OpIn    = "IN"
	// OpNotIn = "NOT_IN"
}

func Compare(left interface{}, operator string, right interface{}) (result bool, err error) {
	// TODO: ...

	operatorFn, ok := operatorMap[operator]
	if !ok {
		return false, fmt.Errorf("invalid operator %v", operator)
	}

	result, err = operatorFn(left, right)
	return
}

func OperatorEqual(left interface{}, right interface{}) (result bool, err error) {

	return
}

func OperatorDiff(left interface{}, right interface{}) (result bool, err error) {
	if left == nil && right == nil {
		return true, nil
	}

	return
}
