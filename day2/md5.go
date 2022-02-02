package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	text := "rina@123rina"
	fmt.Println(createHash(text))

}
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
