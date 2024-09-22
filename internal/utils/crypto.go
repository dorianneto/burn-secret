package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

type cipherData struct {
	Code  string
	Nonce []byte
}

func NewCipherData(code string, nonce []byte) *cipherData {
	return &cipherData{Code: code, Nonce: nonce}
}

func EncryptIt(data string) (*cipherData, error) {
	block, err := aes.NewCipher([]byte(os.Getenv("CRYPTO_KEY")))
	if err != nil {
		return &cipherData{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return &cipherData{}, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return &cipherData{}, err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(data), nil)

	return NewCipherData(base64.StdEncoding.EncodeToString(ciphertext), nonce), nil
}

func DecryptIt(data *cipherData) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(data.Code)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(os.Getenv("CRYPTO_KEY")))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, data.Nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
