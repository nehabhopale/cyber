package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
)

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {

	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func main() {
	fmt.Println("enter text u want to encrypt")
	var originalText string
	fmt.Scanln(&originalText)
	// originalText := "nehabhopaleqwererttt"
	fmt.Println("original text is", originalText)
	mytext := []byte(originalText)
	fmt.Println("enter key  ")
	var keyText string
	fmt.Scanln(&keyText)
	key := []byte(keyText)
	if len(key) > 8 {
		key = key[:8]
	}
	iv := []byte("43218765")

	cryptoText, _ := DesEncryption(key, iv, mytext)
	fmt.Println("cipher text is ", base64.URLEncoding.EncodeToString(cryptoText))
	decryptedText, _ := DesDecryption(key, iv, cryptoText)
	fmt.Println("text after dexcryption is ", string(decryptedText))

}
