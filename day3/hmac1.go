package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func checkIfMsgIsAuthenticated(msgReceived string, hmacGiven string) bool {
	secret := "neha"
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(msgReceived))
	hmacObtained := hex.EncodeToString(h.Sum(nil))
	if hmacObtained == hmacGiven {
		return true
	}
	return false
}
func main() {
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publickey := privatekey.PublicKey
	msg := "some data"
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publickey, []byte(msg), nil)
	secret := "neha"
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(cipher))
	hmacOutput := hex.EncodeToString(h.Sum(nil))
	cipherToSend := hmacOutput + string(cipher)
	fmt.Println(cipherToSend)
	fmt.Println(checkIfMsgIsAuthenticated(cipherToSend[64:], cipherToSend[:64]))
	if checkIfMsgIsAuthenticated(cipherToSend[64:], cipherToSend[:64]) {
		plainText, _ := privatekey.Decrypt(nil, []byte(cipherToSend[64:]), &rsa.OAEPOptions{Hash: crypto.SHA512})
		fmt.Println("Plain text ", string(plainText))
	} else {
		fmt.Println("msg has been modified in between")
	}
}
