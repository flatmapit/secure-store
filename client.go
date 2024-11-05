package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateKey(size int) (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// Encrypt The message must be no longer than the length of the public modulus minus twice the hash length, minus a further 2.
func Encrypt(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func Decrypt(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, message, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
