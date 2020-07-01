package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenToken(t *testing.T) {
	opts := &TokenOpts{
		KeyId:     "1",
		TokenType: ApiKeyToken.String(),
		Duration:  3,
		Unit:      Month,
	}

	tk, err := GenToken(opts)

	assert.Nil(t, err)
	t.Log(tk)
	t.Log(err)
}

func TestIsValidToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXlfaWQiOiIxIiwidHlwIjoiYXBpX2tleSIsImV4cCI6MTY0ODA1NzQ4NiwiaWF0IjoxNTkzNjI1NDg2LCJpc3MiOiJ0ZXN0In0.KkOdv0O48ida_qNITfT4RiHYAX3KYw_rxE2uDTXBWEM`
	// token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.a.v`

	ok := IsValidToken(token)
	t.Log(ok)
}
