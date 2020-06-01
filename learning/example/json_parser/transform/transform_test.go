package transform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrTransformCopy(t *testing.T) {
	p := `[{"operation": "copy", "spec": {"output": "input"}}]`
	i := `{"input":72}`
	out, err := StrTransform(p, i)

	t.Log(out, " | ", err)
}

func TestStrTransformFE(t *testing.T) {
	p := `[
      {
		 "operation": "func_extract",
		 "spec": {
			"out_1": "$DEMO(\"input\", 1, true)",
			"out_2": "$DEMO(\"obj.key\")",
			"out_3": "$DEMO(\"arr\")",
			"out_4": "$DEMO(\"arr_obj[*].key\")",
			"obj_5.key": "$DEMO(\"arr_obj[*].key\")",
			"obj_6.key_normal": "arr_obj[*].key",
			"obj_7": "arr_obj[*]"
		 }
      }
	]`

	i := `{
		"input": "1",
		"obj": {
			"key": 1
		},
		"arr": [1, 2, 3, null],
		"arr_obj": [{"key": 1}, {"key": "2"}, {"key": 3}, {"key": ""}, {"key": null}]
	}`
	out, err := StrTransform(p, i)

	t.Log(out, " | ", err)
}

func TestSkipEmpty(t *testing.T) {
	p := `[
      {
		 "operation": "func_extract",
		 "spec": {
			"out_1": "input",
			"out_11": "input1",

			"out_2": "obj.key",
			"out_22": "obj.key2",
			"out_23": "obj.key3",

			"out_3": "arr",
			"out_4": "arr_obj[*].key",
			"obj_5.key": "arr_obj[*].key",
			"obj_6": "arr_obj[*]"
		 }
      }
	]`

	i := `{
		"input": "",
		"input1": "ok",
		"obj": {
			"key": null,
			"key2": "",
			"key3": "ok"
		},
		"arr": ["1", "", 3, null, true, {"k": 1}],
		"arr_obj": [{"key": 1}, {"key": "2"}, {"key": 3}, {"key": ""}, {"key": null}]
	}`
	out, err := StrTransform(p, i)
	assert.Equal(t, len(`{"out_1":"","out_11":"ok","out_23":"ok","obj_5":{"key":[1,"2",3]},"out_2":null,"out_22":"","out_3":["1",3,true,{"k":1}],"out_4":[1,"2",3],"obj_6":[{"key":1},{"key":"2"},{"key":3},{"key":""},{"key":null}]}`), len(out))
	t.Log(out, " | ", err)
}
