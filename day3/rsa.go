package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

//mess encrypted with private key can be opened with public key
func main() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := privateKey.PublicKey

	msg := "nehaBhopale"
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publicKey, []byte(msg), nil) //ls will encrypt mess using bob n,e
	fmt.Println("encrypted text", string(cipher))

	plaintext, _ := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512}) //bob decrypting using its own private key
	fmt.Println(string(plaintext))
}
