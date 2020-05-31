package operation

import (
	"github.com/qntfy/jsonparser"
	"github.com/qntfy/kazaam/transform"
)

func Copy(spec *transform.Config, data []byte) ([]byte, error) {
	// the internal `Spec` will contain a mapping of source and target keys
	for targetField, sourceFieldInt := range *spec.Spec {
		sourceField := sourceFieldInt.(string)
		// Note: jsonparser.Get() strips quotes from returned strings, so a real
		// transform would need handling for that. We use a Number below for simplicity.
		result, _, _, _ := jsonparser.Get(data, sourceField)
		data, _ = jsonparser.Set(data, result, targetField)
	}
	return data, nil
}
