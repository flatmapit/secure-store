package client

import (
	"crypto/rsa"
	"log"
	"math/rand"
	"testing"
	"time"
)

const rsaLen = 4096

func FuzzEncryptDecrypt(f *testing.F) {
	testcases := []string{"Test plaintext.", "Test case", "", "encrypt this"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(encryptDecrypt)
}

func TestLongPlainText(t *testing.T) {
	const RsaLen = 4096
	charset := "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plainText := []byte(stringWithCharset(1000, charset))

	privateKey, err := GenerateKey(rsaLen)
	if err != nil {
		t.Error(err)
	}

	publicKey := &privateKey.PublicKey

	maxLen := publicKey.N.BitLen() - 2*256 - 2
	log.Println(maxLen)
	_, err = Encrypt(plainText, publicKey)
	if err != rsa.ErrMessageTooLong {
		t.Errorf("Expected error: %s. Error: %s.", rsa.ErrMessageTooLong, err)
	}
}

func encryptDecrypt(t *testing.T, plainTextString string) {
	plainText := []byte(plainTextString)

	privateKey, err := GenerateKey(rsaLen)
	if err != nil {
		t.Error(err)
	}

	publicKey := &privateKey.PublicKey
	cipherText, err := Encrypt(plainText, publicKey)
	if err != nil {
		t.Error(err)
	}

	decryptedText, err := Decrypt(cipherText, privateKey)
	if string(plainText) != string(decryptedText) {
		t.Errorf(`Plaintext: "%s", Decrypted plaintext: "%s"`, plainText, decryptedText)
	}
}

func stringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}
