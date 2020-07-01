package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	opts := &TokenOpts{
		KeyId:     "1",
		TokenType: ApiKeyToken.String(),
		Duration:  3,
		// Unit:      Month,
		Unit: Second,
	}

	tk, err := GenToken(opts)

	assert.Nil(t, err)
	t.Log(tk)
	t.Log(err)
}

func TestIsValidToken(t *testing.T) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXlfaWQiOiIxIiwidHlwIjoiYXBpX2tleSIsImV4cCI6MTY0ODA1NzQ4NiwiaWF0IjoxNTkzNjI1NDg2LCJpc3MiOiJ0ZXN0In0.KkOdv0O48ida_qNITfT4RiHYAX3KYw_rxE2uDTXBWEM`
	// token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.a.v`
	token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJrZXlfaWQiOiIxIiwidHlwIjoiYXBpX2tleSIsImV4cCI6MTU5MzYyODM3MCwiaWF0IjoxNTkzNjI4MzY3LCJpc3MiOiJ0ZXN0In0.nyaGIfHdWisTTuxNEdJ99i2nTGnN8TqhWK-rh_EUJxk`

	ok := IsValidToken(token)
	t.Log(ok)
}

func TestTimeUnit_FromNow(t *testing.T) {
	d, _ := Month.FromNow(1)
	d, _ = Day.FromNow(1)
	d, _ = Day.FromNow(0)
	d, _ = Day.FromNow(2)
	d, _ = Week.FromNow(2)
	d, _ = Month.FromNow(0)
	d, _ = Month.FromNow(1)
	d, _ = Month.FromNow(2)

	t.Log(d)
	t.Log(time.Unix(d, 0))
}
