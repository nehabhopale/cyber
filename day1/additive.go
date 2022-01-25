package main

// import "fmt"

// func main() {
// 	all_letters := "abcdefghijklmnopqrstuvwxyz"
// 	// m := make(map[string]int)
// 	key := 4
// 	// all_letters[(i+key)%len(all_letters)]
// 	for i, item := range all_letters {
// 		fmt.Println(item)
// 		fmt.Println((i + key) % len(all_letters))
// 		// m[item.(string)]=all_letters[(i+key)%len(all_letters)]

// 	}
// }

import (
	"fmt"
	"strings"
)

func removeSpaces(text string) string {
	return strings.ReplaceAll(text, " ", "")
}

func encrypt(plainText string, key int) string {
	input_text := removeSpaces(plainText)
	cipherText := ""
	for _, ltr := range input_text {
		asciiLtr := int(ltr)
		modifiedLtr := string((asciiLtr + key) % 128)
		cipherText = cipherText + modifiedLtr
	}
	return cipherText
}

func decrypt(cipherText string, key int) string {
	plainText := ""
	for _, ltr := range cipherText {
		asciiLtr := int(ltr)
		modifiedLtr := string((asciiLtr - key) % 128)
		plainText = plainText + modifiedLtr
	}
	return plainText
}

func main() {
	plainText := "Neha Bhopale"
	key := 4

	cipherText := encrypt(plainText, key)

	fmt.Println(cipherText)

	plainText1 := decrypt(cipherText, key)

	fmt.Println(plainText1)
}
