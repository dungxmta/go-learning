package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	count := 1
	for i := 0; i < count; i++ {
		// fmt.Printf("%v ", i)
		raw := "tiếng việt nha !@#"
		b := []byte(raw)
		fmt.Println("\nraw:", raw)
		fmt.Println("byte:", b)

		strArr := GoByte2JsonArray(b)
		fmt.Println("strArr:", strArr)

		bArr, err := JsonArray2GoByte(strArr)
		if err != nil {
			t.Fatal(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, b, bArr)
		fmt.Println("->", string(bArr))
	}
}
