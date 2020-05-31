package transform

import "testing"

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
			"obj_5.key": "$DEMO(\"arr_obj[*].key\")"
		 }
      }
	]`

	i := `{
		"input": "1",
		"obj": {
			"key": 1
		},
		"arr": [1, 2, 3],
		"arr_obj": [{"key": 1}, {"key": "2"}, {"key": 3}]
	}`
	out, err := StrTransform(p, i)

	t.Log(out, " | ", err)
}
