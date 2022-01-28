package main

//integrity is given mess is right or wrong
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//encrypt messs ...on encrypted obtain hash
//e:h to bob
//bob hash e mathvh it will h
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
	actualhash := string(data[:32])
	cipherText := data[32:]

	obtainedHash := createHash(string(cipherText))

	if actualhash == (obtainedHash) {
		fmt.Println("hash is right")
		key := []byte(createHash((passphrase)))
		data := cipherText
		block, _ := aes.NewCipher(key)
		gcm, _ := cipher.NewGCM((block))
		nonceSize := gcm.NonceSize()
		nonce := data[:nonceSize]
		cipherText := data[nonceSize:]
		plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
		return plainText
	}
	return nil

	// fmt.Println("hsh", string(actualhash))

	// key := []byte(createHash((passphrase)))
	// block, _ := aes.NewCipher(key)
	// gcm, _ := cipher.NewGCM((block))
	// nonceSize := gcm.NonceSize()
	// nonce := data[:nonceSize]
	// cipherText := data[nonceSize:]
	// plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	// return plainText
}

func main() {
	cipherText := encrypt([]byte("Neha"), "hello")
	fmt.Println("cipher text", string(cipherText))

	hashOfCipher := createHash(string(cipherText))
	cipherWithHash := ""
	cipherWithHash += hashOfCipher
	fmt.Println("actual hash", cipherWithHash)
	cipherWithHash += string(cipherText)
	fmt.Println(cipherWithHash)
	result := decrypt([]byte(cipherWithHash), "hello")

	fmt.Println(string(result))

}
