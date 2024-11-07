package security

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

// Encrypt Encrypts message using OAEP with SHA-256
func Encrypt(message []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}

func MaxEncryptBitLen(publicKey *rsa.PublicKey) int {
	return publicKey.N.BitLen() - 2*256 - 2
}

func Decrypt(message []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, message, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
