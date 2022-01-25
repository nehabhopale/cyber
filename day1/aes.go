package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash((passphrase)))
	block, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM((block))
	nonceSize := gcm.NonceSize()
	nonce := data[:nonceSize]
	cipherText := data[nonceSize:]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}
func main() {
	cipherText := encrypt([]byte("Neha"), "hello")
	fmt.Println(string(cipherText))
	fmt.Println(string(decrypt(cipherText, "hello")))
}
