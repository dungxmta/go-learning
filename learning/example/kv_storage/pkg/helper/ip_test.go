package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextIPStr(t *testing.T) {
	inps := map[string]string{
		"1.0.0.0":   "1.0.0.1",
		"1.0.0.2":   "1.0.0.3",
		"1.0.0.3":   "1.0.0.4",
		"1.0.1.0":   "1.0.1.1",
		"1.0.0.255": "1.0.1.0",
	}

	for inp, expected := range inps {
		v := NextIPStr(inp)
		assert.Equal(t, expected, v)
	}
}
