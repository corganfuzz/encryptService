package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type EncryptServiceInstance struct{}

var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

func (EncryptServiceInstance) Decrypt(_ context.Context, key string, text string) (string, error) {

	if key == "" || text == "" {
		return "", errEmpty
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	ciphertext, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}

var errEmpty = errors.New("Secret Key or Text should not be empty")
