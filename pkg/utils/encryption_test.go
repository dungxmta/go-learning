package utils

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

const secretKeyTest = "super_secret_key_any_abc"

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
