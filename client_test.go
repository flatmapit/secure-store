package client

import (
	"testing"
)

func TestCrypto(t *testing.T) {
	privateKey, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	plainText := []byte("Test plaintext.")

	cipherText, err := Encrypt(plainText, publicKey)
	if err != nil {
		panic(err)
	}

	decryptedText, err := Decrypt(cipherText, privateKey)
	if string(plainText) != string(decryptedText) {
		t.Fatalf(`Plaintext: "%s", Decrypted plaintext: "%s"`, plainText, decryptedText)
	}
}
