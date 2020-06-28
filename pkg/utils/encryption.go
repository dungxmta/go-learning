package utils

/*
Remember not to confuse encryption and decryption with hashing.
When you encrypt something, you’re anticipating being able to get that data back.

When you’re hashing data using something like bcrypt,
	you’re anticipating never being able to read the hashed value again,
	but instead compare against the hashed value.
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
)

// *** NOTE: secret_key ***
//
// should be the AES key,
// either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

//
// Wrap encrypt/decrypt function to return hex string instead of bytes slice of unicode string
//

func EncryptHex(key, plaintext string) (string, error) {
	ciphertext, err := Encrypt([]byte(key), []byte(plaintext))
	if err != nil {
		return "", err
	}

	// encode unicode byte to hex string
	encrypted := hex.EncodeToString(ciphertext)
	return encrypted, nil
}

func DecryptHex(key, hexEncrypted string) (string, error) {
	// decode hex string back to unicode byte
	encrypted, err := hex.DecodeString(hexEncrypted)
	if err != nil {
		return "", err
	}

	decrypted, err := Decrypt([]byte(key), encrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

//
// @src https://itnext.io/encrypt-data-with-a-password-in-go-b5366384e291
//

// encrypt the "data" that can be decrypted using the "key"
//
// @return something like []byte("ZÆ.��e�^mΣ���t�")
func Encrypt(key, data []byte) ([]byte, error) {
	// initializing the block cipher based on the key using Advanced Encryption Standard
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// the block cipher only encrypts the first 16 bytes of data -> wrap cipher.Block -> called "modes"
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	// using Galois Counter Mode with a standard nonce length -> generate a randomized nonce
	//
	// the "nonce" stands for: number once used,
	//  and it's a piece of data that should not be repeated
	//  and only used once in combination with any particular key.
	//  meaning: don't repeat the combination of a key and a nonce more than once.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	// encrypted data
	//  nonce doesn't have to be secret, it just has to be unique
	//  prepended "nonce" to "data" -> nonce+data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	// split "nonce" and "data" <- nonce+data
	if len(data) < gcm.NonceSize() {
		return nil, errors.New("invalid data")
	}
	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]
	// nonce[0] = 1

	// using GCM to decrypting
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
