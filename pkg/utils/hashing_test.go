package utils

import "testing"

func TestHash(t *testing.T) {
	raw := `api_key_random_string`

	var lst []string

	for i := 0; i < 5; i++ {
		hashed, err := Hash(raw)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(hashed)
		lst = append(lst, hashed)
	}

	for _, hashed := range lst {
		equal := CompareHash(hashed, raw)
		if !equal {
			t.Fail()
		}
	}

	v := CompareHash(lst[0], raw+" ")
	t.Log(v)
}
