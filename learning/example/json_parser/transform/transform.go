package transform

import (
	"github.com/qntfy/kazaam"
	"log"
	"testProject/learning/example/json_parser/transform/operation"
)

// use the default config to have access to built-in kazaam transforms
var kc = kazaam.NewDefaultConfig()

func init() {
	NewConfig()
}

func NewConfig() {
	// register the new custom transform called "copy" which supports copying the
	// value of a top-level key to another top-level key
	kc.RegisterTransform("copy", operation.Copy)
	err := kc.RegisterTransform("func_extract", operation.FuncExtract)
	if err != nil {
		log.Fatal(err)
	}
}

func StrTransform(pipeline string, input string) (string, error) {
	// k, _ := kazaam.New(`[{"operation": "copy", "spec": {"output": "input"}}]`, kc)
	// kazaamOut, err := k.TransformJSONStringToString(`{"input":72}`)
	k, _ := kazaam.New(pipeline, kc)
	out, err := k.TransformJSONStringToString(input)
	return out, err
}
