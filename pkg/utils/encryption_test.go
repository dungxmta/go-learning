package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const secretKeyTest = "5a9c73b64b221fa7a9e90425c175631c24fcacb5f57f44f2a34174cb45f9567d70a21ccfdb420560f24ed40100862823dd3f0a8cf7e9c0b0d5bfa38bad"

func TestGenerateKey(t *testing.T) {
	k, err := GenerateKey()
	t.Log(k)
	t.Log(string(k))

	h := hex.EncodeToString(k)
	t.Log(h)
	t.Log(err)
}

func TestEncrypt_Bytes(t *testing.T) {
	raw := "!@# chuỗi tiếng việt à nha $%^"
	rawByte := []byte(raw)
	// t.Log(rawByte)

	// encrypt
	ciphertext, err := Encrypt([]byte(secretKeyTest), rawByte)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ciphertext)
	t.Log(string(ciphertext))

	// encrypted to string
	strEncrypted := hex.EncodeToString(ciphertext)
	t.Log(strEncrypted)

	encrypted, err := hex.DecodeString(strEncrypted)
	if err != nil {
		t.Fatal(err)
	}

	// decrypted, err := Decrypt([]byte(secretKeyTest), ciphertext)
	decrypted, err := Decrypt([]byte(secretKeyTest), encrypted)
	// test := "super_secret_ke1"
	// decrypted, err := Decrypt([]byte(test), encrypted)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rawByte, decrypted)
	assert.Equal(t, raw, string(decrypted))
}

func TestEncryptHex(t *testing.T) {
	raw := "!@# chuỗi tiếng việt à nha $%^"

	encrypted, err := EncryptHex(secretKeyTest, raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)

	decrypted, err := DecryptHex(secretKeyTest, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, raw, decrypted)
}

func TestTimeEncrypt_01_OneByOne(t *testing.T) {
	raw := "!@# chuỗi tiếng việt à nha $%^!@# chuỗi tiếng việt à nha $%^!@# chuỗi tiếng việt à nha $%^"

	limit := 10

	for i := 0; i < limit; i++ {
		s := fmt.Sprintf("%v_%v", raw, i)

		encrypted, err := EncryptHex(secretKeyTest, s)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(encrypted)
	}
}

func TestTimeEncrypt_02_WithChan(t *testing.T) {
	raw := "!@# chuỗi tiếng việt à nha $%^!@# chuỗi tiếng việt à nha $%^!@# chuỗi tiếng việt à nha $%^"

	type Out struct {
		Id      int
		Encrypt string
		Err     error
	}
	var wg sync.WaitGroup

	limit := 10
	ch := make(chan Out, limit)

	for i := 0; i < limit; i++ {
		wg.Add(1)
		s := fmt.Sprintf("%v_%v", raw, i)

		go func(id int, queueCh chan<- Out) {
			defer wg.Done()

			encrypted, err := EncryptHex(secretKeyTest, s)

			queueCh <- Out{
				Id:      id,
				Encrypt: encrypted,
				Err:     err,
			}
		}(i, ch)
	}

	wg.Wait()
	close(ch)

	for v := range ch {
		t.Log(v.Id)
		if v.Err != nil {
			t.Fatal(v.Err)
		}
		t.Log(v.Encrypt)
	}
}

func BenchmarkEncryptHex(b *testing.B) {
	// b.StopTimer()
	// rBytes, _ := GenerateKey()
	// b.StartTimer()
	raw := "5a9c73b64b221fa7a9e90425c175631c24fcacb5f57f44f2a34174cb45f95"

	for i := 0; i < b.N; i++ {
		// _, err := EncryptHex(secretKeyTest, string(rBytes))
		_, err := EncryptHex(secretKeyTest, raw)
		_, err = EncryptHex(secretKeyTest, raw)
		_, err = EncryptHex(secretKeyTest, raw)
		if err != nil {
			b.Fatal(err)
		}
	}
}
